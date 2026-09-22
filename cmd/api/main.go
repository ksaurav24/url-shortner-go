package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ksaurav24/url-shortner-go/internal/config"
	"github.com/ksaurav24/url-shortner-go/internal/db"
	"github.com/ksaurav24/url-shortner-go/internal/handlers"
)

func main() {

	cfg := config.MustLoad()
	db.Connect(cfg.DatabaseURL)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", handlers.Healthz)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("Starting the server")
	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("Failed to start the server. Error: %v", err)
	}
}
