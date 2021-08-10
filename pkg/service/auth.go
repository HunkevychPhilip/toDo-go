package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/HunkevychPhilip/todo/pkg/repository"
	"github.com/HunkevychPhilip/todo/pkg/types"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt         = "gasdjfh780ap98whg9apsdg89adsf98g"
	signingKey   = "asfjsdf8sa9fdsa98f8s9adfjiojiiojoj"
	tokenExpTime = 1 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (a *AuthService) CreateUser(user *types.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return a.repo.CreateUser(user)
}

func (a *AuthService) GenerateToken(data *types.SignInData) (string, error) {
	user, err := a.repo.GetUser(data.Username, generatePasswordHash(data.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &types.CustomJwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (a *AuthService) ParseToken(str string) (int, error) {
	token, err := jwt.ParseWithClaims(str, &types.CustomJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*types.CustomJwtClaims)
	if !ok {
		return 0, errors.New("invalid claims type")
	}

	return claims.UserID, nil
}

func generatePasswordHash(pwd string) string {
	hash := sha1.New()
	hash.Write([]byte(pwd))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
