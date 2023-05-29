package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func EmailOTPSender() {

	url := "https://api.brevo.com/v3/smtp/email"

	payload := strings.NewReader(`{"to":[{"email":"devadiga.rohit@gmail.com","name":"95110"}],"templateId":1}`)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("api-key", "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")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
