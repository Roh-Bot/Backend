package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	data := []byte("Hi There, World!")

	hash := HashData(data)
	fmt.Printf("Hashed Text: %s\n", hash)

	// Verify if a given data matches the hash
	match := VerifyHash(data, hash)
	fmt.Println("Match:", match)
}

func HashData(data []byte) string {
	// Create a new SHA256 hash instance
	hash := sha256.New()

	// Write the data to the hash
	hash.Write(data)

	// Get the hashed result
	hashedData := hash.Sum(nil)

	// Convert the hashed data to hexadecimal string
	hashedText := hex.EncodeToString(hashedData)

	return hashedText
}

func VerifyHash(data []byte, hash string) bool {
	// Convert the hash string back to a byte slice
	hashedData, _ := hex.DecodeString(hash)

	// Create a new SHA256 hash instance
	verifyHash := sha256.New()

	// Write the data to the hash
	verifyHash.Write(data)

	// Get the hashed result for the given data
	computedHash := verifyHash.Sum(nil)

	// Compare the hashes
	return bytes.Equal(hashedData, computedHash)
}
