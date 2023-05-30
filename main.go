package main

import (
	"github.com/Roh-Bot/Backend/handlers"
)

func main() {
	//currentTime := time.Now().Unix()
	//user_id := 112
	//plaintext := []byte(fmt.Sprintf("%d$$%d", currentTime, user_id))
	//cipherText := make([]byte, aes.BlockSize)
	//fmt.Println(len(cipherText))
	//
	//cipherTextString := (fmt.Sprintf("%s", utils.AESEncryption(plaintext, cipherText)))
	//fmt.Println("CipherTExtSTring:", cipherTextString)
	//
	//newCipherText := make([]byte, 16)
	//copy(newCipherText, []byte(cipherTextString))
	//fmt.Println("NewCipherText", len(newCipherText))
	//fmt.Printf("%s", utils.AESDecryption(newCipherText))

	handlers.Start()
}
