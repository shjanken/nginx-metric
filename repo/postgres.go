package repo

import (
	"fmt"

	metric "com.github.shjanken/nginx-metric"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

const DB_ERROR = `database connect is nil.`

type postgres struct {
	db *pg.DB
}

// InsertError is a error with current log
type InsertError struct {
	Log    metric.Log
	ErrMsg string
}

func (ie *InsertError) Error() string {
	return fmt.Sprintf("insert log %v error: %v", ie.Log, ie.ErrMsg)
}

func (pr *postgres) Insert(logs []metric.Log) error {
	if pr.db == nil {
		panic(DB_ERROR)
	} // 如果 postgres 的数据库连接是 Nil，说明这个是一个bug

	_, err := pr.db.Model(&logs).Insert()

	if err != nil {
		return fmt.Errorf("insert logs failuire. %w. ", err)
	}

	return nil
}

func (pr *postgres) createSchema() error {
	err := pr.db.Model((*metric.Log)(nil)).CreateTable(&orm.CreateTableOptions{
		Temp: true,
	})
	if err != nil {
		return err
	}
	return nil
}
