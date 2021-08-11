package service

import (
	"github.com/HunkevychPhilip/todo/pkg/repository"
	"github.com/HunkevychPhilip/todo/pkg/types"
)

type Authorization interface {
	CreateUser(*types.User) (int, error)
	GenerateToken(*types.SignInData) (string, error)
	ParseToken(string) (int, error)
}

type List interface {
	Create(int, *types.List) (int, error)
	Get(int, int) (*types.List, error)
	GetAll(int) ([]*types.List, error)
	Delete(int, int) error
}

type Item interface {
}

type Service struct {
	Authorization
	List
	Item
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		List:          NewListService(repos.List),
	}
}
