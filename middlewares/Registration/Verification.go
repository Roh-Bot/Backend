package Registration

import (
	"context"
	"fmt"
	"github.com/Roh-Bot/Backend/middlewares/SessionHandling"
	"github.com/Roh-Bot/Backend/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Verification(c echo.Context) error {
	if time.Now().Unix()-utils.CurrentTime > 900 {
		return c.String(http.StatusOK, "Time Expired")
	}
	pool := utils.PostgresConnectionPool()
	_, err := pool.Query(context.Background(), `UPDATE users SET is_email_verified=true WHERE email=$1`, Register.Email)
	if err != nil {
		fmt.Println("Query unsuccessful:", err)
	} else {
		fmt.Println("Query Successful")
	}
	SessionHandling.Signin(c, Register.User_Id)
	fmt.Println(Register.User_Id)
	return c.Redirect(http.StatusFound, "http://localhost:8000/") // ConfigChange := url in config

}

//UPDATE users SET is_email_verified = true WHERE email='Dhebug@God.com
