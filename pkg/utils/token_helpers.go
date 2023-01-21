package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type GenerateTokenParams struct {
	Username  string
	AccountId string
	Superuser bool
}

type ValidateTokenResult struct {
	Username   string
	AccountId  string
	Authorized bool
	Superuser  bool
}

func GetJwtSecretKey() ([]byte, error) {
	key, err := GetParameter("JwtSecretKey")
	if err != nil {
		return []byte(""), err
	}

	jwtSecretKey := []byte(key)

	return jwtSecretKey, nil
}

// Users can only get an auth token if they've already verified their email
func GenerateToken(args GenerateTokenParams) (string, error) {
	// Create a JWT
	token := jwt.New(jwt.SigningMethodHS256)

	jwtSecretKey, err := GetJwtSecretKey()
	if err != nil {
		return "", err
	}

	// Modify the JWT with registered claims
	// https://auth0.com/docs/secure/tokens/json-web-tokens/json-web-token-claims
	claims := token.Claims.(jwt.MapClaims)

	// Stardard registered claims
	claims["sub"] = args.Username                           // Subject of the JWT (user)
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix() // Expiration time
	claims["iat"] = time.Now().Unix()                       // Issued at time
	claims["nbf"] = time.Now().Unix()                       // Not before time

	// Custom claims
	claims["act"] = args.AccountId
	claims["auth"] = true // Auth token can be used to access secure resources
	claims["admin"] = args.Superuser

	// Sign the JWT with a secret key
	tokenString, err := token.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return "", fmt.Errorf("generating JWT Token failed: %w", err)
	}

	return tokenString, nil
}

func GenerateRefreshToken(args GenerateTokenParams) (string, error) {
	// Create a JWT
	token := jwt.New(jwt.SigningMethodHS256)

	jwtSecretKey, err := GetJwtSecretKey()
	if err != nil {
		return "", err
	}

	// Modify the JWT with registered claims
	// https://auth0.com/docs/secure/tokens/json-web-tokens/json-web-token-claims
	claims := token.Claims.(jwt.MapClaims)

	// Stardard registered claims
	claims["sub"] = args.Username                         // Subject of the JWT (user)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Expiration time
	claims["iat"] = time.Now().Unix()                     // Issued at time
	claims["nbf"] = time.Now().Unix()                     // Not before time

	// Custom claims
	claims["act"] = args.AccountId
	claims["auth"] = false // Refresh token cannot be used to access secure resources
	claims["admin"] = args.Superuser

	// Sign the JWT with a secret key
	tokenString, err := token.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return "", fmt.Errorf("generating JWT refresh Token failed: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(token string) (ValidateTokenResult, error) {
	emptyResult := ValidateTokenResult{}

	jwtSecretKey, err := GetJwtSecretKey()
	if err != nil {
		return emptyResult, err
	}

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return emptyResult, fmt.Errorf("invalidate token: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return emptyResult, fmt.Errorf("invalid token claim")
	}

	result := ValidateTokenResult{
		Username:   claims["sub"].(string),
		AccountId:  claims["act"].(string),
		Authorized: claims["auth"].(bool),
		Superuser:  claims["admin"].(bool),
	}

	return result, nil
}
