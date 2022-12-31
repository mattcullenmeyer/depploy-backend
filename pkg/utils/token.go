package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretJWTKey = []byte("gosecretkey") // TODO: store as env variable

func GenerateToken(username string) (string, error) {
	// Create a JWT
	token := jwt.New(jwt.SigningMethodHS256)

	// Modify the JWT with registered claims
	// https://auth0.com/docs/secure/tokens/json-web-tokens/json-web-token-claims
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = username                                // Subject of the JWT (user)
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix() // Expiration time
	claims["iat"] = time.Now().Unix()                       // Issued at time
	claims["nbf"] = time.Now().Unix()                       // Not before time

	// Sign the JWT with a secret key
	tokenString, err := token.SignedString([]byte(secretJWTKey))

	if err != nil {
		return "", fmt.Errorf("generating JWT Token failed: %w", err)
	}

	return tokenString, nil
}
