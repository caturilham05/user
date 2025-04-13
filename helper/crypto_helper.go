package helper

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateSecretKey() string {
	// Generate 32 bytes of random data
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	// Encode to Base64
	encodedKey := base64.StdEncoding.EncodeToString(key)
	return encodedKey
}
