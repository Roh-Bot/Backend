package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(password string) string {
	var bytePassword = []byte(password)
	hashedPassword := sha256.Sum256(bytePassword)
	stringHash := hex.EncodeToString(hashedPassword[:])
	return stringHash
}
