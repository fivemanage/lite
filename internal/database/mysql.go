package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type MySQL struct{}

func (r *MySQL) Connect() *bun.DB {
	sqldb, err := sql.Open("mysql", "root:pass@localhost/test")
	if err != nil {
		panic(err)
	}

	return bun.NewDB(sqldb, mysqldialect.New())
}
