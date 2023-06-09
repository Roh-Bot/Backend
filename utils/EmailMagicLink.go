package utils

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var CurrentTime = time.Now().Unix()

func EmailMagicLink(user_id int, email string) {
	url := "https://api.brevo.com/v3/smtp/email"

	plaintext := []byte(fmt.Sprintf("%d$$%d", CurrentTime, user_id))
	cipherText := make([]byte, aes.BlockSize)
	encryptedText := AESEncryption(plaintext, cipherText)

	hashQuery := base64.URLEncoding.EncodeToString(encryptedText)
	fmt.Println(hashQuery)

	payload := strings.NewReader(fmt.Sprintf(`{"sender":{"id":1},"to":[{"email":"%s"}],
"replyTo":{"email":"rohit.devadiga@geneticminds.com","name":"Rohit"},
"htmlContent":"<!DOCTYPE html> <html> <body> <a href = 'localhost:8000/verifylink?hash=%s&user_id=%d'>Please click the link below to verify your account </a> </body> </html>",
"textContent":"Verify your otp","subject":"Verification"}`, email, hashQuery, user_id))
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("api-key", "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")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	//fmt.Println(res)
	fmt.Println(string(body))

}
