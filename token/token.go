package token

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/imega/teleport-server/config"
)

var (
	RsaPublicKey  []byte
	RsaPrivateKey []byte
)

type Claims struct {
	jwt.StandardClaims
}

func init() {
	key, err := config.GetConfigValue("RSA_PUBLIC_KEY")
	if err != nil {
		log.Fatalf("failed to read public key %s", err)
	}
	RsaPublicKey = []byte(key)

	key, err = config.GetConfigValue("RSA_PRIVATE_KEY")
	if err != nil {
		log.Fatalf("failed to read private key %s", err)
	}
	RsaPrivateKey = []byte(key)
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
