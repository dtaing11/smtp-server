package auth

import (
	"crypto/rand"
	"encoding"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func passwordHashing(password string) (string, error) {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashPassword), err
}

func generateAPIKey(lenght int) (string, error) {

	byte := make([]byte, lenght)

	if _, err := rand.Read(byte); err != nil {
		return "", err
	}
	encoded := base64.URLEncoding.EncodeToString(byte)

	return encoded, nil
}
