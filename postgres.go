package metric

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

const dbError = `database connect is nil.`

type postgres struct {
	db *pg.DB
}

// NewPostgre return a repo use postgresql backend
func NewPostgre(dbname, user, password string) Repo {
	return &postgres{
		db: pg.Connect(&pg.Options{
			User:     user,
			Password: password,
			Database: dbname,
		}),
	}
}

func (pr *postgres) Insert(logs []Log) error {
	if pr.db == nil {
		panic(dbError)
	} // 如果 postgres 的数据库连接是 Nil，说明这个是一个bug

	_, err := pr.db.Model(&logs).Insert()

	if err != nil {
		return fmt.Errorf("insert logs failuire. %w. ", err)
	}

	return nil
}

func (pr *postgres) Close() error {
	return pr.db.Close()
}

func (pr *postgres) createSchema() error {
	err := pr.db.Model((*Log)(nil)).CreateTable(&orm.CreateTableOptions{
		Temp: true,
	})
	if err != nil {
		return err
	}
	return nil
}
