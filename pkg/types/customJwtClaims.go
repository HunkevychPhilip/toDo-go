package types

import "github.com/dgrijalva/jwt-go"

type CustomJwtClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}
