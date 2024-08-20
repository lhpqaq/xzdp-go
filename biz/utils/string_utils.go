package utils

import (
	"math/rand"
	"regexp"
	"time"
)

func ValidateMobile(phone string) bool {
	var re = regexp.MustCompile(`^1[3-9]\d{9}$`)
	return re.MatchString(phone)
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}
