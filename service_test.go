package metric

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"com.github.shjanken/nginx-metric/test"
	"github.com/satyrius/gonx"
)

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

// read the data from gnox.Entry
// if gnox.Entry.Field() return error, return ""
func readDataFromGnoxEntry(entry *gonx.Entry, field string) string {
	val, err := entry.Field(field)
	if err != nil {
		return ""
	}
	return val
}

// test read the data from source
func TestRead(t *testing.T) {
	c := strings.NewReader(test.TestConfig)
	d := strings.NewReader(test.TestLogs)
	p := provider{
		ch:     make(chan Item, 10),
		data:   d,
		config: c,
	}
	service := NewService(&p, nil)

	var logs []Item
	var err error
	if logs, err = service.Read(); err != nil {
		t.Fatalf("%+v\n", err)
	}

	if count := len(logs); count != 6 {
		t.Fatalf("except logs count 6, but hava %d", count)
	}
}

func TestSave(t *testing.T) {

}
