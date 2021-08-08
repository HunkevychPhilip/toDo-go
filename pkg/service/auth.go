package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/HunkevychPhilip/todo/pkg/repository"
	"github.com/HunkevychPhilip/todo/pkg/types"
)

const salt = "gasdjfh780ap98whg9apsdg89adsf98g"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (a *AuthService) CreateUser(user *types.User) (int, error) {
	user.Password = a.generatePasswordHash(user.Password)

	return a.repo.CreateUser(user)
}

func (a *AuthService) generatePasswordHash(pwd string) string {
	hash := sha1.New()
	hash.Write([]byte(pwd))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
