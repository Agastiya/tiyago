package jwt

import (
	"fmt"
	"time"

	"github.com/agastiya/tiyago/dto"
	"github.com/agastiya/tiyago/pkg/constant"
	"github.com/golang-jwt/jwt/v5"
)

type IJwt interface {
	GenerateToken(account dto.LoginResponse, key string) (newToken string, err error)
	VerifyToken(tokenString string, key string) (jwt.MapClaims, error)
}

func (j Jwt) MapKeyWithEnv(key string) JwtConfig {

	mapKey := map[string]JwtConfig{
		"secret_key": {
			Key: j.JwtPackage.SecretKey,
			Exp: constant.ONEHOUR,
		},
		"refresh_secret_key": {
			Key: j.JwtPackage.RefreshSecretKey,
			Exp: constant.THREEHOURS,
		},
	}

	return mapKey[key]
}

func (j Jwt) GenerateToken(account dto.LoginResponse, key string) (string, error) {

	if key == "" {
		return "", fmt.Errorf("invalid key")
	}

	secretKey := j.MapKeyWithEnv(key)
	if secretKey.Key == "" {
		return "", fmt.Errorf("key not found!")
	}

	claims := jwt.MapClaims{
		"id":       account.Id,
		"fullname": account.Fullname,
		"username": account.Username,
		"email":    account.Email,
		"exp":      time.Now().Add(time.Duration(secretKey.Exp) * time.Second).Unix(),
		"iat":      time.Now().Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(secretKey.Key))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (j Jwt) VerifyToken(tokenString string, key string) (jwt.MapClaims, error) {

	if key == "" {
		return nil, fmt.Errorf("invalid key")
	}

	secretKey := j.MapKeyWithEnv(key)
	if secretKey.Key == "" {
		return nil, fmt.Errorf("key not found!")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secretKey.Key), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
