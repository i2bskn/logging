package logging

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRequestID(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}
