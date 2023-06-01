package Login

import (
	"fmt"
	"github.com/Roh-Bot/Backend/models"
	"github.com/Roh-Bot/Backend/models/Login"
	"github.com/Roh-Bot/Backend/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

func LoginController(c echo.Context) error {
	var login Login.LoginStruct

	var reponseFunc models.Response
	var response Login.LoginResponse
	if err := c.Bind(&login); err != nil {
		fmt.Println("Binding Error")
		return c.NoContent(http.StatusBadRequest)
	}
	v := validator.New()
	if err := v.Struct(login); err != nil {
		fmt.Println("Validation Error")
		responseJSON := reponseFunc.ResponseWriter(200, map[string]string{"code": "002", "message": "Invalid Credentails"}, map[string]string{})
		return c.JSON(http.StatusBadRequest, responseJSON)
	}
	loginPass := utils.SHA256(login.Password)

	loginStatus, loginMessage, loginUserId := utils.CheckLoginCredentials(login.Email, login.Mobile, loginPass)
	if !loginStatus {
		response = Login.LoginResponse{
			StatusCode: 400,
			Error: map[string]string{
				"code":    "002",
				"message": "Invalid Credentails",
			},
			Data: map[string]string{},
		}
		return c.String(400, fmt.Sprintf(`%+v\n%s`, response, loginMessage))
	}
	fmt.Println(loginUserId)

	response = Login.LoginResponse{
		StatusCode: 200,
		Error:      map[string]string{},
		Data: map[string]string{
			"user_id": "01",
		},
	}
	return c.JSON(http.StatusOK, fmt.Sprintf(`%+v%s`, response, loginMessage))
}
