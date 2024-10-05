package migrate

import "github.com/spf13/cobra"

var (
	RootCmd = &cobra.Command{
		Use:   "db",
		Short: "Database migrations",
	}

	InitCmd = &cobra.Command{
		Use:   "init",
		Short: "Create migration tables",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	MigrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Migrate database",
	}
)
