package main

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env. Assuming environment variables are set directly.")
	}

	dbURL := os.Getenv("MY_DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set.")
	}

	migrationsPath := os.Getenv("MIGRATIONS_PATH")
	if migrationsPath == "" {
		migrationsPath = "./database/migrations"
		log.Printf("MIGRATIONS_PATH environment variable not set, using default: %s", migrationsPath)
	}
	if !strings.HasPrefix(migrationsPath, "file://") {
		migrationsPath = "file://" + migrationsPath
	}

	migration, err := migrate.New(migrationsPath, dbURL)
	defer gracefulShutDown(migration)
	if err != nil {
		log.Fatal(err)
	}

	if err := upOrDownMigration(migration); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	} else if errors.Is(err, migrate.ErrNoChange) {
		log.Println("No new migrations to apply.")
		return
	}

	log.Println("Migrations applied successfully!")
}

func gracefulShutDown(migration *migrate.Migrate) {
	err, _ := migration.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migration closed")
}

func upOrDownMigration(migrate *migrate.Migrate) error {
	if len(os.Args) < 2 || strings.ToLower(os.Args[1]) == "up" {
		return migrate.Up()
	} else if len(os.Args) > 1 && strings.ToLower(os.Args[1]) == "down" {
		return migrate.Steps(-1)
	}
	return errors.New("unrecognized command")
}
