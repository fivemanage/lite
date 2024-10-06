package migrations

import (
	"context"
	"fmt"

	"github.com/fivemanage/lite/internal/database"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")
		db.RegisterModel((*database.User)(nil))
		_, err := db.NewCreateTable().Model((*database.User)(nil)).Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	}, nil)
}
