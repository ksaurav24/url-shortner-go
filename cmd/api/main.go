package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/ksaurav24/url-shortner-go/internal/config"
	"github.com/ksaurav24/url-shortner-go/internal/db"
	"github.com/ksaurav24/url-shortner-go/internal/handlers"
)

func main() {

	cfg := config.MustLoad()
	db, db_err := db.Connect(cfg.DatabaseURL)
	if db_err != nil {
		log.Fatalf("Failed to connect DB: %v", db_err)
	}
	mux := http.NewServeMux()

	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})

	logger := slog.New(logHandler)

	uh := handlers.NewUrlHandler(db, logger)

	mux.HandleFunc("GET /healthz", handlers.Healthz)
	mux.HandleFunc("GET /url", uh.GetUrls)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("Starting the server")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start the server. Error: %v", err)
	}
}
