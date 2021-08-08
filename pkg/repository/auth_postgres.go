package repository

import (
	"fmt"
	"github.com/HunkevychPhilip/todo/pkg/types"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (a *AuthPostgres) CreateUser(user *types.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password) values($1, $2, $3) RETURNING id;", userTable)
	row := a.db.QueryRow(
		query,
		user.Username,
		user.Username,
		user.Password,
	)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
