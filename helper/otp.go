package helper

import (
	"crypto/rand"
	"math/big"
)

func OTPGenerator(length int) (string, error) {
	const charset = "0123456789"

	result := make([]byte, length)

	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}

		result[i] = charset[num.Int64()]
	}

	return string(result), nil
}
