package postgres

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
	"path"
)

type Migrator struct {
	m *migrate.Migrate
}

func NewMigrator(db *sql.DB, dbName string) (*Migrator, error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	migrationsDir := path.Join(cwd, "internal/message_approval/infrastructure/persistence/postgres/migrations")
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file:%s", migrationsDir), dbName, driver)
	if err != nil {
		return nil, err
	}

	return &Migrator{m: m}, nil
}

func (m *Migrator) MigrateUp() error {
	return m.m.Up()
}

func (m *Migrator) MigrateDown() error {
	return m.m.Down()
}
