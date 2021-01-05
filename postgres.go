package metric

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

const dbError = `database connect is nil.`

// PostgresRepo impletement the repo interface.
// use postgres db as backend
type PostgresRepo struct {
	db *pg.DB
}

// NewPostgre return a repo use postgresql backend
func NewPostgre(dbname, user, password string) *PostgresRepo {
	return &PostgresRepo{
		db: pg.Connect(&pg.Options{
			User:     user,
			Password: password,
			Database: dbname,
		}),
	}
}

// Insert the logs to db
func (pr *PostgresRepo) Insert(logs []Log) error {
	if pr.db == nil {
		panic(dbError)
	} // 如果 postgres 的数据库连接是 Nil，说明这个是一个bug

	_, err := pr.db.Model(&logs).Insert()

	if err != nil {
		return fmt.Errorf("fileprovider insert failure %w. ", err)
	}

	return nil
}

// Close db connect
func (pr *PostgresRepo) Close() error {
	return pr.db.Close()
}

// CreateSchema 创建需要的表
func (pr *PostgresRepo) CreateSchema() error {
	return pr.db.Model((*Log)(nil)).CreateTable(&orm.CreateTableOptions{
		Temp: false,
	})
}

// DropSchema delete table
func (pr *PostgresRepo) DropSchema() error {
	return pr.db.Model((*Log)(nil)).DropTable(&orm.DropTableOptions{
		IfExists: true,
		Cascade:  true,
	})
}
