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

		Convey("should panic when postgre db backend is nil", func() {
			panicPG := postgres{}

			So(func() {
				panicPG.Insert([]metric.Log{
					{Request: "fake request"},
				})
			}, ShouldPanicWith, DB_ERROR)
		})

		Convey("should success insert data into database", func() {
			defer fakePG.db.Close() // 测试完成以后关闭数据库连接

			logs := []metric.Log{
				{Request: "test request 1"},
				{Request: "test request 2"},
				{Request: "test request 3"},
			}

			if err := fakePG.createSchema(); err != nil {
				t.Fatalf("create table failure, %v", err)
				t.Fail()
			}

			err := fakePG.Insert(logs)

			So(err, ShouldBeNil)
		})
	})
}
