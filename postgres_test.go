package metric

import (
	"testing"
	"time"

	"github.com/go-pg/pg/v10"
	. "github.com/smartystreets/goconvey/convey"
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

func TestPostgresRepo(t *testing.T) {
	Convey("test postgres repository", t, func() {
		fakePG := PostgresRepo{
			db: pg.Connect(&pg.Options{
				User:     "postgres",
				Password: "postgres",
				Database: "metric",
			}),
		}
		defer fakePG.db.Close()

		Convey("should panic when postgre db backend is nil", func() {
			panicPG := PostgresRepo{}

			So(func() {
				panicPG.Insert([]Log{
					{Request: "fake request"},
				})
			}, ShouldPanicWith, dbError)
		})

		Convey("should success insert data into database", func() {
			logs := []Log{
				{Request: "test request 1"},
				{Request: "test request 2"},
				{Request: "test request 3"},
			}

			if err := fakePG.DropSchema(); err != nil {
				t.Fatalf("drop table failure %v", err)
			}

			if err := fakePG.CreateSchema(); err != nil {
				t.Fatalf("create table failure, %v", err)
			}

			err := fakePG.Insert(logs)
			// select the data from database
			count, err := fakePG.db.Model((*Log)(nil)).Count()
			var inserted []Log
			err = fakePG.db.Model(&inserted).Select()

			So(err, ShouldBeNil)
			So(count, ShouldEqual, 3)
			for _, l := range inserted {
				So(l, ShouldBeIn, logs)
			}
		})
	})
}

func TestNewPostgre(t *testing.T) {
	Convey("test NewPostgre function", t, func() {
		Convey("should return a object", func() {
			postgre := NewPostgre("metric", "postgre", "postgres")
			So(postgre, ShouldNotBeNil)
		})
	})
}
