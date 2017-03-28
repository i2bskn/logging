package logging

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRequestID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	b[6] = (b[6] & 0x0f) | 0x40 // version
	b[8] = (b[8] & 0xbf) | 0x80 // variant

	uuid := make([]byte, 36)
	hex.Encode(uuid[0:8], b[0:4])
	uuid[8] = '-'
	hex.Encode(uuid[9:13], b[4:6])
	uuid[13] = '-'
	hex.Encode(uuid[14:18], b[6:8])
	uuid[18] = '-'
	hex.Encode(uuid[19:23], b[8:10])
	uuid[23] = '-'
	hex.Encode(uuid[24:], b[10:])
	return string(uuid)
}
