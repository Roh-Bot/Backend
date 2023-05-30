package utils

import (
	"crypto/aes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var CurrentTime = time.Now().Unix()
var user_id = 112

func EmailMagicLink() {
	url := "https://api.brevo.com/v3/smtp/email"

	plaintext := []byte(fmt.Sprintf("%d$$%d", CurrentTime, user_id))
	cipherText := make([]byte, aes.BlockSize)
	AESEncryption(plaintext, cipherText)
	//AESDecryption(cipherText)

	fmt.Println()
	payload := strings.NewReader(fmt.Sprintf(`{"sender":{"id":1},"to":[{"email":"devadiga.rohit@gmail.com"}],
"replyTo":{"email":"rohit.devadiga@geneticminds.com","name":"Rohit"},
"htmlContent":"<!DOCTYPE html> <html> <body> <a href = 'localhost:8080/verifylink?hash=%x&user_id=%d'>Please click the link below to verify your account </a> </body> </html>",
"textContent":"Verify your otp","subject":"Verification"}`, AESEncryption(plaintext, cipherText), user_id))
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("api-key", "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")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
