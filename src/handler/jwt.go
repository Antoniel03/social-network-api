package handler

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

type CustomClaims struct {
	Name     string `json:"name"`
	Lastname string `json:"Lastname"`
	Id       int    `json:"id"`
	jwt.StandardClaims
}

func GenerateJWT(signingKey string, claims CustomClaims) (*string, error) {
	claims.ExpiresAt = time.Now().Unix() + 60
	claims.Issuer = "test"
	claims.IssuedAt = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return nil, err
	}
	return &ss, nil
}

func validateHeaderAuth(authType string, credential string) error {
	if authType == "" || credential == "" {
		return errors.New("Malformed credentials")
	}
	if authType != "Bearer" {
		return errors.New("Invalid authorization type")
	}
	return nil
}

func validateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("guacamole"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ValidateAuthorizedRequest(credentials string) error {
	if credentials == "" {
		return errors.New("No credentials provided")
	}
	splitCredentials := strings.Split(credentials, " ")
	if err := validateHeaderAuth(splitCredentials[0], splitCredentials[1]); err != nil {
		return err
	}

	if _, err := validateToken(splitCredentials[1]); err != nil {
		return err
	}
	return nil
}
