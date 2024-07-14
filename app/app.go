package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"url-shortener/apis"
	. "url-shortener/daos"
	"url-shortener/database/sqlite"
	"url-shortener/services"
)

const (
	shutdownTimeoutS = 3
)

func StartApp() {
	// configure db
	dao := NewDao(sqlite.Open(), sqlite.OpenNonConcurrent())
	defer dao.Destroy()
	dao.CreateTables()

	// configure dependencies
	service := services.NewService(dao)
	router := apis.InitRouter(service)

	// start
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.Handler(),
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Server shutdown...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(shutdownTimeoutS)*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v\n", err)
	}

	<-ctx.Done()
	log.Println("Server closed")
}
