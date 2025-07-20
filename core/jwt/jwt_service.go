package jwtservice

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("mi_clave_secreta_super_segura")


type JWTClaims struct {
	UsuarioID int    `json:"usuario_id"`
	Email     string `json:"email"`
	jwt.RegisteredClaims
}

// Crear token JWT
func GenerateJWT(usuarioID int, email string) (string, error) {
	claims := JWTClaims{
		UsuarioID: usuarioID,
		Email:     email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Expira en 24h
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// Verificar token JWT
func ValidateJWT(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}