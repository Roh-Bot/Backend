package main

import "github.com/Roh-Bot/Backend/handlers"

func main() {
	//currentTime := time.Now().Unix()
	//user_id := 112
	//plaintext := []byte(fmt.Sprintf("%d$$%d", currentTime, user_id))
	//cipherText := make([]byte, aes.BlockSize)
	//fmt.Printf("%x\n", utils.AESEncryption(plaintext, cipherText))
	//fmt.Printf("%s", utils.AESDecryption(cipherText))
	handlers.Start()
}
