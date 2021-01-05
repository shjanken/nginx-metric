package metric

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReadData(t *testing.T) {
	Convey("test read data from simple data provider", t, func() {

		Convey("should read 100 lines", func() {
			simpleProvider := &simpleProvider{
				data: createFakeData(100),
			}
			service := NewService(simpleProvider, nil)

			items, err := service.Read()
			if err != nil {
				t.Fatalf("read data failure %v", err)
			}

			So(items, ShouldNotBeNil)
			So(len(items), ShouldEqual, 100)
			for _, v := range items {
				So(simpleProvider.data, ShouldContain, v)
			}
		})

		Convey("should success read 1000 lines", func() {
			simpleProvider := &simpleProvider{
				data: createFakeData(1000),
			}
			service := NewService(simpleProvider, nil)

			items, err := service.Read()
			if err != nil {
				t.Fatalf("read data failure, %v", err)
			}

			So(items, ShouldNotBeNil)
			So(len(items), ShouldEqual, 1000)
			for _, v := range items {
				So(simpleProvider.data, ShouldContain, v)
			}
		})

		SkipConvey("should success read 2000 lines", func() {
			simpleProvider := &simpleProvider{
				data: createFakeData(2000),
			}
			service := NewService(simpleProvider, nil)

			items, err := service.Read()
			if err != nil {
				t.Fatalf("read data failure, %v", err)
			}

			So(items, ShouldNotBeNil)
			So(len(items), ShouldEqual, 2000)
			for _, v := range items {
				So(simpleProvider.data, ShouldContain, v)
			}
		})
	})
}

func TestSave(t *testing.T) {
	Convey("test service save function", t, func() {
		Convey("should success save the data", func() {
			fakeProvider := simpleProvider{
				data: createFakeData(1234),
			}
			fakeRepo := simpleRepo{}
			service := NewService(&fakeProvider, &fakeRepo)

			if err := service.Save(); err != nil {
				t.Fatalf("service save data failure %v", err)
			}

			// fmt.Println(len(fakeProvider.data))
			for _, v := range fakeProvider.data {
				So(fakeRepo.data, ShouldContain, v.Log)
			}
			So(len(fakeRepo.data), ShouldEqual, 1234)
		})
	})
}

type simpleProvider struct {
	data []*Item
}

func (sp *simpleProvider) ReadData(ch chan *Item) error {
	for _, v := range sp.data {
		ch <- v
	}
	close(ch)
	return nil
}

type simpleRepo struct {
	data []Log
}

func (repo *simpleRepo) Insert(logs []Log) error {
	for _, v := range logs {
		repo.data = append(repo.data, v)
	}
	return nil
}

func (repo *simpleRepo) Close() error {
	return nil
}

func createFakeData(count int) (items []*Item) {
	for i := 0; i < count; i++ {
		item := &Item{
			Log: Log{
				Request: fmt.Sprintf("fake request %d", i),
			},
			Error: nil,
		}
		items = append(items, item)
	}
	return
}
