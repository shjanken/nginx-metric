package metric

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReadData(t *testing.T) {
	Convey("test read data from simple data provider", t, func() {
		simpleProvider := &simpleProvider{}

		Convey("should read 100 lines", func() {
			service := NewService(simpleProvider, nil)

			items, err := service.Read()
			if err != nil {
				t.Fatalf("read data failure %v", err)
			}

			So(items, ShouldNotBeNil)
			So(len(items), ShouldEqual, 100)
			for _, v := range items {
				So(simpleProvider.data, ShouldContain, v.Log.Request)
			}
		})
	})
}

type simpleProvider struct {
	data []string
}

func (sp *simpleProvider) ReadData(ch chan *Item) {
	for i := 0; i < 100; i++ {
		item := &Item{
			Log: Log{
				Request: fmt.Sprintf("fake request %d", i),
			},
			Error: nil,
		}
		ch <- item
		sp.data = append(sp.data, item.Log.Request)
	}
	close(ch)
}
