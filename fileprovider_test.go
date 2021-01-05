package metric

import (
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	TestConfig = `
http {
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent"'

}
`
)

func TestSaveToDB(t *testing.T) {
	Convey("test save data to db", t, func() {
		Convey("should success save data", func() {
			pg := NewPostgre("metric", "postgres", "postgres")
			defer pg.Close()

			pg.DropSchema()
			pg.CreateSchema()

			logs := []Log{
				{
					Request:     "fake request 1",
					TimeLocal:   time.Now(),
					HTTPReferer: "no referer",
				},
				{
					Request:     "fake request 2",
					TimeLocal:   time.Now(),
					HTTPReferer: "no referer",
				},
			}

			err := pg.Insert(logs)
			if err != nil {
				t.Fatalf("save data to db failure %v", err)
			}

			So(err, ShouldBeNil)
		})
	})
}

func TestReadDataFromFile(t *testing.T) {
	if os.Getenv("BIGFILE") == "" {
		t.Skip("skip read testdata/access.log file test")
	}

	Convey("should read 5000 lines from access.log file", t, func() {
		file, err := os.Open("testdata/access.log")
		if err != nil {
			t.Fatalf("read accesslog file failure.%v", err)
		}
		defer file.Close()
		fileProvider := fileProvider{
			logFile: file,
			config:  TestConfig,
		}
		ch := make(chan *Item, 10)
		var items []*Item

		go fileProvider.ReadData(ch)
		for item := range ch {
			items = append(items, item)
		}

		So(items, ShouldNotBeNil)
		So(len(items), ShouldEqual, 1018448)
	})
}
