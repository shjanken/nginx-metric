package metric

import (
	"os"
	"testing"

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

func TestReadDataFromFile(t *testing.T) {
	Convey("should read 500 lines from access.log file", t, func() {
		file, err := os.Open("testdata/access.500.log")
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
		So(len(items), ShouldEqual, 500)
	})
}
