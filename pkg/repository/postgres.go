package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const (
	segmentsTable             = "segments"
	usersSegmentsTable        = "users_segments"
	usersSegmentsHistoryTable = "users_segments_history"
)

type DbConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	Ssl      string
	Driver   string
}

func ConnectToDb(cfg DbConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.Driver, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.Ssl))
	if err != nil {
		log.Fatalf("aboba cfg:%v\n", cfg)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
