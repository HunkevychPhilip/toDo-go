package repository

import (
	"github.com/HunkevychPhilip/todo/pkg/types"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user *types.User) (id int, err error)
	GetUser(username, pwd string) (user *types.User, err error)
}

type List interface {
	Create(int, *types.List) (int, error)
	Get(int, int) (*types.List, error)
	GetAll(int) ([]*types.List, error)
	Delete(int, int) error
}

type Item interface {
}

type Repository struct {
	Authorization
	List
	Item
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		List:          NewListPostgres(db),
	}
}
