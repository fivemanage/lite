package database

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:user"`
	ID            int    `bun:"id,pk,autoincrement"`
	Name          string `bun:"name"`
	Email         string `bun:"email"`
	PasswordHash  string `bun:"password_hash"`
	AuthProvider  string `bun:"auth_provider"`
	AuthID        string `bun:"auth_id"`
	Avatar        string `bun:"avatar"`
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
