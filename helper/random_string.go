package helper

import (
	"math/rand"
	"strings"
	"time"
)

var seedRandom = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	sb.Grow(length)

	for i := 0; i < length; i++ {
		randomCharacter := charset[seedRandom.Intn(len(charset))]
		sb.WriteByte(randomCharacter)
	}

	return sb.String()
}
