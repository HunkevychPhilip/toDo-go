package service

import (
	"github.com/HunkevychPhilip/todo/pkg/repository"
	"github.com/HunkevychPhilip/todo/pkg/types"
)

type Authorization interface {
	CreateUser(user *types.User) (id int, err error)
	GenerateToken(data *types.SignInData) (token string, err error)
	ParseToken(str string) (id int, err error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
