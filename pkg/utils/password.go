package utils

import "golang.org/x/crypto/bcrypt"

func VerifyPassword(hashedPassword string, userPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword))
}
