package main

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
	"github.com/miladvatankhah/maker-checker/configs"
	pg "github.com/miladvatankhah/maker-checker/internal/message_approval/infrastructure/persistence/postgres"
	"github.com/miladvatankhah/maker-checker/pkg/clients/postgres"
	"log"
)

func init() {
	if err := godotenv.Load(".env.dev"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load DB config: %v", err)
	}

	postgresClient, err := postgres.NewPostgresClient(cfg.Postgres)
	if err != nil {
		log.Fatalf("Failed to initialize PostgreSQL client: %v", err)
	}
	defer postgresClient.Close()

	m, err := pg.NewMigrator(postgresClient.DB, cfg.Postgres.DBName)
	if err != nil {
		log.Fatalf("Failed to initialize Migrator: %v", err)
	}

	if err = m.MigrateUp(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Failed to run the migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}
