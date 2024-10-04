package database

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

type SQLite struct{}

func (r *SQLite) Connect() *bun.DB {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared") // fivemanage.db
	if err != nil {
		panic(err)
	}

	return bun.NewDB(sqldb, sqlitedialect.New())
}
