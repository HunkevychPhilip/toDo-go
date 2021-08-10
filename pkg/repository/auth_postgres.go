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

func (a *AuthPostgres) GetUser(username, pwd string) (*types.User, error) {
	var user types.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2;", userTable)
	if err := a.db.Get(&user, query, username, pwd); err != nil {
		return nil, err
	}

	return &user, nil
}
