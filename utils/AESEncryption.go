package utils

import (
	"crypto/aes"
	"fmt"
	"github.com/zenazn/pkcs7pad"
)

var key = []byte("1234567890123456")
var cipher, _ = aes.NewCipher(key)

func AESEncryption(data, cipherText []byte) []byte {
	paddedData := pkcs7pad.Pad(data, aes.BlockSize)
	cipher.Encrypt(cipherText, paddedData)
	fmt.Printf("CipherText: %x\n", cipherText)
	return cipherText
}

func AESDecryption(cipherText []byte) []byte {
	if len(cipherText) < aes.BlockSize {
		fmt.Println("Invalid cipherText length")
		return []byte("No")
	}

	decrypted := make([]byte, aes.BlockSize)
	cipher.Decrypt(decrypted, cipherText)
	unpaddedData, err := pkcs7pad.Unpad(decrypted)
	if err != nil {
		fmt.Println("Padding error", err)
	}

	fmt.Printf("DecryptedText: %s\n", unpaddedData)
	return unpaddedData
}

//func PKCS7Pad(data []byte) []byte {
//	blockSize := aes.BlockSize
//	padding := blockSize - (len(data) % blockSize)
//	padText := bytes.Repeat([]byte{byte(padding)}, padding)
//	return append(data, padText...)
//}
//
//func PKCS7Unpad(data []byte) []byte {
//	length := len(data)
//	unpadding := int(data[length-1])
//	fmt.Println(data)
//	fmt.Println(unpadding)
//	if unpadding > length {
//		fmt.Println("Invalid padding")
//		return nil
//	}
//	return data[:(length - unpadding)]
//}
