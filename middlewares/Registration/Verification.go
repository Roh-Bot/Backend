package Registration

import (
	context2 "context"
	"fmt"
	"github.com/Roh-Bot/Backend/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Verification(context echo.Context) error {
	if time.Now().Unix()-utils.CurrentTime > 900 {
		return context.String(http.StatusOK, "Time Expired")
	}
	pool := utils.PostgresConnectionPool()
	_, err := pool.Query(context2.Background(), `UPDATE users SET is_email_verified=true WHERE email=$1`, register.Email)
	if err != nil {
		fmt.Println("Query unsuccessful:", err)
	} else {
		fmt.Println("Query Successful")
	}

	return context.Redirect(http.StatusFound, "http://localhost:8080/")

}

//UPDATE users SET is_email_verified = true WHERE email='Dhebug@God.com
