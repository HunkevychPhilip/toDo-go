package repository

import (
	"errors"
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

func (l *ListPostgres) Get(userID, listID int) (*types.List, error) {
	var list types.List

	getListQuery := fmt.Sprintf(
		"SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2;",
		todoListsTable, usersListsTable)

	if err := l.db.Get(&list, getListQuery, userID, listID); err != nil {
		return nil, err
	}

	return &list, nil
}

func (l *ListPostgres) GetAll(userID int) ([]*types.List, error) {
	var lists []*types.List

	getListsQuery := fmt.Sprintf(
		"SELECT tl.id, tl.title, tl.description FROM %s ul INNER JOIN %s tl on tl.id = ul.list_id WHERE ul.user_id = $1;",
		usersListsTable, todoListsTable)
	if err := l.db.Select(&lists, getListsQuery, userID); err != nil {
		return nil, err

	}

	return lists, nil
}

func (l *ListPostgres) Delete(userID, listID int) error {
	deleteListQuery := fmt.Sprintf(
		"DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2;",
		todoListsTable, usersListsTable)

	res, err := l.db.Exec(deleteListQuery, userID, listID)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("no such list")
	}

	return nil
}
