package domain

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	Session string `json:"session"`
	jwt.RegisteredClaims
}
