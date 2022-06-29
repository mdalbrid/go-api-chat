package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	chatsTable    = "chats"
	messagesTable = "messages"
	usersTable    = "users"
)

type Config struct {
	User     string
	Password string
	DbName   string
	Host     string
	Port     string
}

// NewPostgresDB ...
func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable ",
			cfg.User, cfg.Password, cfg.DbName, cfg.Host, cfg.Port))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
