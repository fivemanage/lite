package main

import (
	"context"
	"fmt"
	"log"
	nethttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fivemanage/lite/internal/database"
	"github.com/fivemanage/lite/internal/http"
	"github.com/fivemanage/lite/internal/service/authservice"
	"github.com/fivemanage/lite/migrate"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "fivemanage",
		Short: "Open-source, easy-to-use file hosting service.",
	}

	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run fivemanage application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)

			port := viper.GetInt("port")
			driver := viper.GetString("driver")

			err := godotenv.Load()
			// TODO: Only for development
			if err != nil {
				log.Fatal("Error loading .env file.")
			}

			db := database.New(driver, "")
			store := db.Connect()

			authService := authservice.New(store)
			server := http.NewServer(authService)

			srv := &nethttp.Server{
				Addr:    fmt.Sprintf(":%d", port),
				Handler: server.Engine,
			}

			go func() {
				fmt.Printf("Server is running on port %d...\n", port)
				if err := srv.ListenAndServe(); err != nil && err != nethttp.ErrServerClosed {
					log.Fatalf("listen: %s\n", err)
				}
			}()

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit
			log.Println("Shutdown Server ...")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("Server Shutdown:", err)
			}
			select {
			case <-ctx.Done():
				log.Println("timeout of 5 seconds.")
			}
			log.Println("Server exiting")
		},
	}
)

func init() {
	rootCmd.PersistentFlags().String("driver", "sqlite", "Database driver")
	runCmd.Flags().Int("port", 8080, "Port to serve Fivemanage")

	viper.BindPFlag("driver", rootCmd.PersistentFlags().Lookup("driver"))
	viper.BindPFlag("port", runCmd.Flags().Lookup("port"))

	rootCmd.AddCommand(runCmd)

	rootCmd.AddCommand(migrate.RootCmd)
	migrate.RootCmd.AddCommand(migrate.InitCmd, migrate.MigrateCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
