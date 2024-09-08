package database

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:user"`
	ID            int    `bun:"id,pk,autoincrement"`
	Name          string `bun:"name"`
	Email         string `bun:"email"`
}

type Store interface {
	Connect() *bun.DB
}

func New(option string, dsn string) Store {
	switch option {
	case "mysql":
		return &MySQL{}
	case "sqlite":
		return &SQLite{}
	}

	return nil
}
