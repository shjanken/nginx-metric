package metric

import (
	"fmt"
	"io"
	"strings"
	"testing"

	test "com.github.shjanken/nginx-metric/testdata"
	"github.com/satyrius/gonx"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMetricService(t *testing.T) {
	Convey("test metrics service functions", t, func() {
		c := strings.NewReader(test.TestConfig)
		d := strings.NewReader(test.TestLogs)
		simpleProvider := provider{
			ch:     make(chan Item, 10),
			data:   d,
			config: c,
		}

		Convey("test read function", func() {

			Convey("Should return error", func() {
				errProvider := errorProvider{
					errMsg: "simple error",
				}
				service := NewService(&errProvider, nil)

				_, err := service.Read()

				So(err, ShouldNotBeNil)
				So(fmt.Sprint(err), ShouldContainSubstring, "simple error")
			})

			Convey("The Read funciton should return a list, count is 6", func() {
				service := NewService(&simpleProvider, nil)

				logs, _ := service.Read()

				So(len(logs), ShouldEqual, 6)
			})

			Convey("should panic if data provider is nil", func() {
				service := NewService(nil, nil)

				So(func() {
					_, _ = service.Read()
				}, ShouldPanic)
			})
		})

		Convey("test save funciton", func() {

			// save function need repo
			r := repo{}

			Convey("should return err if provider.ReadData return err", func() {
				errProvider := errorProvider{
					errMsg: "simple error",
				}
				ser := NewService(&errProvider, &r)

				err := ser.Save()

				So(err, ShouldNotBeNil)
				So(fmt.Sprint(err), ShouldEqual, "simple error")
			})

			Convey("should panic if backend repo is nil", func() {
				service := NewService(&simpleProvider, nil)

				So(func() {
					_ = service.Save()
				}, ShouldPanicWith, "the repo backend is nil")
			})

			Convey("should success save the 6 items", func() {
				service := NewService(&simpleProvider, &r)
				service.Save()

				So(len(r.data), ShouldEqual, 6)
				// check the result
				for _, v := range r.data {
					So(test.TestLogs, ShouldContainSubstring, v.Request)
				}
			})
		})
	})
}

// create test data provider
type provider struct {
	ch     chan Item
	data   io.Reader
	config io.Reader
}

// read the data from test data provider
func (p *provider) ReadData() (<-chan Item, error) {
	reader, err := gonx.NewNginxReader(p.data, p.config, "main")
	if err != nil {
		return nil, fmt.Errorf("create nginx reader failure. %w", err)
	}

	// read all the data from reader
	for {
		entry, err := reader.Read()
		if err == io.EOF {
			close(p.ch)
			break
		} else if err != nil {
			return nil, fmt.Errorf("read data from data provider failure. %w", err)
		}

		request := readDataFromGnoxEntry(entry, "request")

		// create the metric.Item struct
		item := Item{
			Log{
				Request: request,
			},
		}

		// send the item to chan
		p.ch <- item
	}

	return p.ch, nil
}

// this provider is used to fire error
type errorProvider struct {
	errMsg string
}

func (ep *errorProvider) ReadData() (<-chan Item, error) {
	return nil, fmt.Errorf(ep.errMsg)
}

type repo struct {
	data []Log
}

func (r *repo) Insert(logs []Log) error {
	r.data = logs
	return nil
}

// read the data from gnox.Entry
// if gnox.Entry.Field() return error, return ""
func readDataFromGnoxEntry(entry *gonx.Entry, field string) string {
	val, err := entry.Field(field)
	if err != nil {
		return ""
	}
	return val
}
