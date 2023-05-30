package Registration

import (
	"fmt"
	"github.com/Roh-Bot/Backend/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Verification(context echo.Context) error {
	if time.Now().Unix()-utils.CurrentTime > 60 {
		return context.String(http.StatusOK, "Time Expired")
	}
	hash := []byte(context.QueryParam("hash"))
	newCipherText := make([]byte, 16)
	copy(newCipherText, hash[:16])
	fmt.Printf("Decrypted:%s", utils.AESDecryption(newCipherText))

	return context.String(http.StatusOK, fmt.Sprintf("%s", newCipherText))

}
