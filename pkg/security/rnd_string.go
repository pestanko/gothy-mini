package security

import "math/rand"

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString generate a random string of lenght
func RandomString(length int) string {
	return RandomStringBytes(length, letters)
}

// RandomStringBytes generate a random string of lenght
func RandomStringBytes(length int, letters string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
