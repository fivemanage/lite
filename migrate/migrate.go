package migrate

import (
	"fmt"
	"strings"

	"github.com/fivemanage/lite/internal/database"
	"github.com/fivemanage/lite/migrate/migrations"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/uptrace/bun/migrate"
)

var (
	RootCmd = &cobra.Command{
		Use:   "db",
		Short: "Database migrations",
	}

	InitCmd = &cobra.Command{
		Use:   "init",
		Short: "Create migration tables",
		Run: func(cmd *cobra.Command, args []string) {
			driver := viper.GetString("driver")
			db := database.New(driver, "").Connect()

			migrator := migrate.NewMigrator(db, migrations.Migrations)
			err := migrator.Init(cmd.Context())
			if err != nil {
				panic(err)
			}
		},
	}

	MigrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Migrate database",
		Run: func(cmd *cobra.Command, args []string) {
			driver := viper.GetString("driver")
			db := database.New(driver, "").Connect()

			migrator := migrate.NewMigrator(db, migrations.Migrations)
			if err := migrator.Lock(cmd.Context()); err != nil {
				panic(err)
			}
			defer migrator.Unlock(cmd.Context()) //nolint:errcheck

			group, err := migrator.Migrate(cmd.Context())
			if err != nil {
				panic(err)
			}
			if group.IsZero() {
				fmt.Printf("there are no new migrations to run (database is up to date)\n")
			}
			fmt.Printf("migrated to %s\n", group)
		},
	}

	UnlockCmd = &cobra.Command{
		Use: "unlock",
		Run: func(cmd *cobra.Command, args []string) {
			driver := viper.GetString("driver")
			db := database.New(driver, "").Connect()

			migrator := migrate.NewMigrator(db, migrations.Migrations)
			err := migrator.Unlock(cmd.Context())
			if err != nil {
				panic(err)
			}
		},
	}

	LockCmd = &cobra.Command{
		Use: "lock",
		Run: func(cmd *cobra.Command, args []string) {
			driver := viper.GetString("driver")
			db := database.New(driver, "").Connect()

			migrator := migrate.NewMigrator(db, migrations.Migrations)
			err := migrator.Lock(cmd.Context())
			if err != nil {
				panic(err)
			}
		},
	}

	CreateMigrationCmd = &cobra.Command{
		Use:   "create",
		Short: "Create database migration",
		Run: func(cmd *cobra.Command, args []string) {
			driver := viper.GetString("driver")
			db := database.New(driver, "").Connect()

			migrator := migrate.NewMigrator(db, migrations.Migrations)

			name := strings.Join(args, "_")
			mf, err := migrator.CreateGoMigration(cmd.Context(), name)
			if err != nil {
				panic(err)
			}
			fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
		},
	}
)
