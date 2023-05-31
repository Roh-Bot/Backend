package main

import "github.com/Roh-Bot/Backend/handlers"

func main() {
	handlers.Start()
}

//currentTime := time.Now().Unix()
//user_id := 112
//plaintext := []byte(fmt.Sprintf("%d$$%d", currentTime, user_id))
//cipherText := make([]byte, aes.BlockSize)
//
//encryptedText := utils.AESEncryption(plaintext, cipherText)
//
//hashQuery := base64.URLEncoding.EncodeToString(encryptedText)
//fmt.Printf("%s\n", hashQuery)
//
//DecodedhashQuery, _ := base64.URLEncoding.DecodeString(hashQuery)
//DecryptedText := utils.AESDecryption([]byte(DecodedhashQuery))
//fmt.Printf("%s", DecryptedText)
