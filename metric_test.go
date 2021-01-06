package metric

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMetricService(t *testing.T) {
	if os.Getenv("BIGFILE") == "" {
		t.Skip("skip test with large access.log file")
	}

	Convey("integrate test", t, func() {
		Convey("should read all lines from access.log file", func() {
			file, err := os.Open("testdata/access.log")
			if err != nil {
				t.Fatalf("read file failure %v", err)
			}
			accessProvider := NewFileProvider(file, testConfig)
			service := NewService(accessProvider, nil)

			items, err := service.Read()
			if err != nil {
				t.Fatalf("service read data failure. %v", err)
			}

			So(len(items), ShouldEqual, 1018448)
		})

		Convey("should success insert access.log data to database", func() {
			// data provider
			file, err := os.Open("testdata/access.log")
			if err != nil {
				t.Fatalf("read file failure %v", err)
			}
			accessProvider := NewFileProvider(file, testConfig)
			//database repo
			pg := NewPostgre("metric", "postgres", "postgres")
			defer pg.Close()

			if err = pg.DropSchema(); err != nil {
				t.Fatalf("drop table error %v", err)
			}

			if err = pg.CreateSchema(); err != nil {
				t.Fatalf("create table error %v", err)
			}
			// service
			service := NewService(accessProvider, pg)

			if err = service.Save(); err != nil {
				t.Fatalf("save data to database failure %v", err)
			}

			So(err, ShouldBeNil)
		})
	})
}

const testConfig = `
http {
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent"'

}
`
