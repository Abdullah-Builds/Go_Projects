package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Abdullah-Builds/Go_Projects/internal/config"
	"github.com/Abdullah-Builds/Go_Projects/internal/http/handlers/student"
	"github.com/Abdullah-Builds/Go_Projects/internal/storage/sqlite"
)

func main() {
	//load config
	cfg := config.MustLoad()

	//database setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	slog.Info("Storage Initialized")

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/student", student.New(storage))
	router.HandleFunc("GET /api/student/{id}", student.GetByID(storage))
	router.HandleFunc("GET /api/student", student.GetAllStudents(storage))

	//setup server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}
	slog.Info("Server Started", slog.String("Addr", cfg.Address))

	block := make(chan os.Signal, 1)
	signal.Notify(block, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Failed to Start Server")
			log.Fatalf("Failed to Start Server: %v", err)
		}
	}()

	<-block

	close(block)
	log.Println("Shutting Down the Server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to Shutdown", slog.String("", err.Error()))
	}

	slog.Info("Server Shutdown")
}
