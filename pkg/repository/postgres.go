package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const PostgresDriverName = "postgres"

const (
	userTable       = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect(PostgresDriverName, cfg.buildConnStr())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (cfg *Config) buildConnStr() string {
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.DBName, cfg.Username, cfg.Password, cfg.SSLMode,
	)
}
