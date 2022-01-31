package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/kokhno-nikolay/lets-go-chat/internal/config"
)

func NewClient(driver string, cfg *config.Config) (*sql.DB, error) {
	pgConnString := fmt.Sprintf(
		`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		cfg.PostgresHostName, cfg.PostgresPort, cfg.PostgresUsername, cfg.PostgresPassword, cfg.PostgresDatabaseName,
	)

	db, err := sql.Open(driver, pgConnString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("postgres: connected to database has been successfully")
	return db, nil
}
