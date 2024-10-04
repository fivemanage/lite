package main

import (
	"context"
	"log"
	nethttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fivemanage/lite/internal/database"
	"github.com/fivemanage/lite/internal/http"
	"github.com/fivemanage/lite/internal/service/authservice"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	// TODO: Only for development
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	db := database.New("sqlite", "")
	store := db.Connect()

	authService := authservice.New(store)

	server := http.NewServer(authService)

	srv := &nethttp.Server{
		Addr:    ":8080",
		Handler: server.Engine,
	}

	go func() {
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
}
