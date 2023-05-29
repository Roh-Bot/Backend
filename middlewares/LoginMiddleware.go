package middlewares

import (
	"fmt"
	"github.com/Roh-Bot/Backend/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

func LoginController(c echo.Context) error {
	var login models.LoginStruct

	var response models.LoginResponse
	if err := c.Bind(&login); err != nil {
		fmt.Println("Binding Error")
		return c.NoContent(http.StatusBadRequest)
	}
	v := validator.New()
	if err := v.Struct(login); err != nil {
		fmt.Println("Validation Error")
		response = models.LoginResponse{
			StatusCode: 400,
			Error: map[string]string{
				"code":    "002",
				"message": "Invalid Credentails",
			},
			Data: map[string]string{},
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	response = models.LoginResponse{
		StatusCode: 200,
		Error:      map[string]string{},
		Data: map[string]string{
			"user_id": "01",
		},
	}
	return c.JSON(http.StatusBadRequest, response)
}
