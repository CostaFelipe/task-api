package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/CostaFelipe/task-api/config"
	_ "github.com/go-sql-driver/mysql"
)

func NewConnection(cfg *config.Config) (*sql.DB, error) {
	dsn := config.DataBaseURL(cfg)

	db, err := sql.Open(cfg.DBDriver, dsn)
	if err != nil {
		return nil, fmt.Errorf("error ao abrir conexão: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("erro de conectar o banco: %w", err)
	}

	return db, nil
}
