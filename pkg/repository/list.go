package repository

import (
	"fmt"
	"github.com/HunkevychPhilip/todo/pkg/types"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ListPostgres struct {
	db *sqlx.DB
}

func NewListPostgres(db *sqlx.DB) *ListPostgres {
	return &ListPostgres{
		db: db,
	}
}

func (l *ListPostgres) Create(userID int, list *types.List) (int, error) {
	tx, err := l.db.Begin()
	if err != nil {
		return 0, err
	}

	var listID int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES($1, $2) RETURNING id;", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	err = row.Scan(&listID)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			logrus.WithError(err).Error("Transaction rollback failed")
		}

		return 0, err
	}

	bindListToUserQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2);", usersListsTable)
	_, err = tx.Exec(bindListToUserQuery, userID, listID)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			logrus.WithError(err).Error("Transaction rollback failed")
		}

		return 0, err
	}

	return listID, tx.Commit()
}
