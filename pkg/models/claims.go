package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	UsuarioID int64 `json:"usuario_id"`
	jwt.StandardClaims
}
