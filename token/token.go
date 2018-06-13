package token

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var (
	RsaPublicKey  []byte
	RsaPrivateKey []byte
)

type Claims struct {
	jwt.StandardClaims
}

// Create создание токена
func Create(ID string, expiresAt int64) (string, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(RsaPrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key, %s", err)
	}

	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Id:        ID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to signet token, %s", err)
	}

	return signedToken, nil
}

// Valid проверка токена
func Valid(signedToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwt.ParseRSAPublicKeyFromPEM(RsaPublicKey)
	})
	if err != nil {
		return nil, fmt.Errorf("token is invalid")
	}

	return token.Claims.(*Claims), nil
}
