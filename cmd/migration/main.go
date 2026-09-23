package main

import (
	"log"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ksaurav24/url-shortner-go/internal/config"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("usage: migrate<up | down>")
	}

	cfg := config.MustLoad()

	m, err := migrate.New(
		"file://internal/migrations",
		cfg.DatabaseURL)

	if err != nil {
		log.Fatalf("Failed connection: %v", err)
	}

	migrateDirection := strings.ToLower(os.Args[1])

	switch migrateDirection {
	case "up":
		if err := m.Steps(1); err != nil {
			log.Fatalf("Failed migration: %v", err)
		}
		log.Println("migration successful")
	case "down":
		if err := m.Steps(-1); err != nil {
			log.Fatalf("Failed migration: %v", err)
		}
		log.Println("migration successful")
	default:
		log.Fatalf("Unknown command: %s\nValid commands: <up | down>", migrateDirection)
	}
}
