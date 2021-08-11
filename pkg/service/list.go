package service

import (
	"github.com/HunkevychPhilip/todo/pkg/repository"
	"github.com/HunkevychPhilip/todo/pkg/types"
)

type ListService struct {
	repo repository.List
}

func NewListService(repo repository.List) *ListService {
	return &ListService{
		repo: repo,
	}
}

func (tds *ListService) Create(userID int, todoList *types.List) (int, error) {
	return tds.repo.Create(userID, todoList)
}

func (tds *ListService) GetUserLists(userID int) ([]*types.List, error) {
	return tds.repo.GetUserLists(userID)
}

func (tds *ListService) GetList(userID, listID int) (*types.List, error) {
	return tds.repo.GetList(userID, listID)
}
