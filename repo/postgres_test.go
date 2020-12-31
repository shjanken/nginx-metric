package repo

import (
	"testing"

	metric "com.github.shjanken/nginx-metric"
	"github.com/go-pg/pg/v10"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPostgresRepo(t *testing.T) {
	Convey("test postgres repository", t, func() {
		fakePG := postgres{
			db: pg.Connect(&pg.Options{
				User:     "postgres",
				Password: "postgres",
				Database: "metric",
			}),
		}
		defer fakePG.db.Close()

		Convey("should panic when postgre db backend is nil", func() {
			panicPG := postgres{}

			So(func() {
				panicPG.Insert([]metric.Log{
					{Request: "fake request"},
				})
			}, ShouldPanicWith, dbError)
		})

		Convey("should success insert data into database", func() {
			logs := []metric.Log{
				{Request: "test request 1"},
				{Request: "test request 2"},
				{Request: "test request 3"},
			}

			if err := fakePG.createSchema(); err != nil {
				t.Fatalf("create table failure, %v", err)
			}

			err := fakePG.Insert(logs)
			// select the data from database
			count, err := fakePG.db.Model((*metric.Log)(nil)).Count()
			var inserted []metric.Log
			err = fakePG.db.Model(&inserted).Select()

			So(err, ShouldBeNil)
			So(count, ShouldEqual, 3)
			for _, l := range inserted {
				So(l, ShouldBeIn, logs)
			}
		})
	})
}
