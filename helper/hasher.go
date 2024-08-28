package helper

import "golang.org/x/crypto/bcrypt"

func Hash(value string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), 14)

	return string(bytes), err
}

func CheckHash(value string, hash string) bool {
	hashBytes := []byte(hash)
	valueBytes := []byte(value)

	err := bcrypt.CompareHashAndPassword(hashBytes, valueBytes)

	return err == nil
}
