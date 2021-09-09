package project

import (
	"math/rand"
)

// GenerateToken generates a random token
func GenerateToken(n int) string {
	var letters = []rune("abcdefghijklmnpqrstuvwxyzABCDEFGHIJKLMNPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
