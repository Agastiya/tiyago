package jwt

import (
	"fmt"
	"time"

	"github.com/agastiya/tiyago/dto"
	"github.com/golang-jwt/jwt/v5"
)

type IJwt interface {
	GenerateToken(account dto.LoginResponse, key string) (newToken string, err error)
	VerifyToken(tokenString string, key string) (jwt.MapClaims, error)
}

func (j Jwt) GenerateToken(account dto.LoginResponse, key string) (string, error) {

	switch key {
	case "secret_key":
		key = j.JwtPackage.SecretKey
	case "refresh_secret_key":
		key = j.JwtPackage.RefresSecretKey
	default:
		return "", fmt.Errorf("key not found!")
	}

	if key == "" {
		return "", fmt.Errorf("key not set")
	}

	claims := jwt.MapClaims{
		"id":       account.Id,
		"fullname": account.Fullname,
		"username": account.Username,
		"email":    account.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (j Jwt) VerifyToken(tokenString string, key string) (jwt.MapClaims, error) {

	switch key {
	case "secret_key":
		key = j.JwtPackage.SecretKey
	case "refresh_secret_key":
		key = j.JwtPackage.RefresSecretKey
	default:
		return nil, fmt.Errorf("key not found!")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
