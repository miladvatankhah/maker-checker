package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type Client struct {
	DB *sql.DB
}

func NewPostgresClient(cfg Config) (*Client, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if err = createDatabaseIfNotExists(db, cfg.DBName); err != nil {
		return nil, err
	}

	db.Close()

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return &Client{DB: db}, nil
}

func (c *Client) Close() {
	if err := c.DB.Close(); err != nil {
		log.Printf("Failed to close database: %v", err)
	}
}

// Function to check if the database exists
func databaseExists(db *sql.DB, dbName string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)`
	err := db.QueryRow(query, dbName).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// Function to create the database if it doesn't exist
func createDatabaseIfNotExists(db *sql.DB, dbName string) error {
	exists, err := databaseExists(db, dbName)
	if err != nil {
		return err
	}

	if !exists {
		// Create the database
		_, err := db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
		if err != nil {
			return fmt.Errorf("error creating database: %w", err)
		}
		log.Printf("Database %s created successfully.\n", dbName)
	}

	return nil
}
