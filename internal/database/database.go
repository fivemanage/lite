package database

import (
	"fmt"

	"github.com/uptrace/bun"
)

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

// TODO: Return error if driver is not supported
func New(driver string, dsn string) Store {
	fmt.Println("option", driver)

	switch driver {
	case "mysql":
		return &MySQL{}
	case "sqlite":
		return &SQLite{}
	default:
		return nil
	}
}
