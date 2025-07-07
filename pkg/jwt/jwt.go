package jwt

import (
	"github.com/golang-jwt/jwt"
)

var (
	JwtVar JwtInterface = &JwtService{}
)

type JwtInterface interface {
	// CreateToken(account auth.LoginData) (newToken string, err error)
	// VerifyToken(tokenString string) (claims jwt.MapClaims, err error)
}

func EncodeToken(claims jwt.MapClaims, secretKey []byte) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(secretKey)
	if err != nil {
		return
	}
	return
}

// func (j JwtService) CreateToken(account auth.LoginData) (newToken string, err error) {
// 	if j.ConfigJwt.SecretKey == "" {
// 		return "", fmt.Errorf("secret key not set")
// 	}

// 	claims := jwt.MapClaims{
// 		"id":       account.Id,
// 		"name":     account.Name,
// 		"username": account.Username,
// 		"email":    account.Email,
// 		"active":   account.Active,
// 		"exp":      time.Now().Add(time.Hour * 24).Unix(),
// 		"iat":      time.Now().Unix(),
// 	}

// 	newToken, err = EncodeToken(claims, []byte(j.ConfigJwt.SecretKey))
// 	if err != nil {
// 		return
// 	}

// 	return
// }

func (j JwtService) VerifyToken(tokenString string) (claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(j.ConfigJwt.SecretKey), nil
	})
	if err != nil {
		return
	}

	if !token.Valid {
		return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorSignatureInvalid)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return
	}

	return
}
