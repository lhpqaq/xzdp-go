package utils

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	once sync.Once
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digitBytes = "0123456789"

func RandInit() {
	rand.New(rand.NewSource(time.Now().UnixNano())) // fix the deprecated method rand.Seed(time.Now().UnixNano())
}

func GenerateDigits(length int) string {
	once.Do(RandInit)
	b := make([]byte, length)
	for i := range b {
		b[i] = digitBytes[rand.Intn(len(digitBytes))]
	}
	return string(b)
}

func GenerateLetters(length int) string {
	once.Do(RandInit)
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GenerateAlphaNumeric(length int) string {
	once.Do(RandInit)
	b := make([]byte, length)
	combinedBytes := letterBytes + digitBytes
	for i := range b {
		b[i] = combinedBytes[rand.Intn(len(combinedBytes))]
	}
	return string(b)
}

func RandomUUID() (string, error) {

	randomBytes := make([]byte, 16)
	combinedBytes := letterBytes + digitBytes
	for i := range randomBytes {
		randomBytes[i] = combinedBytes[rand.Intn(len(combinedBytes))]
	}

	randomBytes[6] &= 0x0f // clear version
	randomBytes[6] |= 0x40 // set version to 4 (random uuid)
	randomBytes[8] &= 0x3f // clear variant
	randomBytes[8] |= 0x80 // set to IETF variant
	return fmt.Sprintf("%x", randomBytes), nil

	// return fmt.Sprintf("%x-%x-%x-%x-%x", randomBytes[0:4], randomBytes[4:6], randomBytes[6:8], randomBytes[8:10], randomBytes[10:]), nil
}

func RandomString(n int) string {
	once.Do(RandInit)
	letters := letterBytes + digitBytes
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
