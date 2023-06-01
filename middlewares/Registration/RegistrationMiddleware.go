package Registration

import (
	"context"
	"fmt"
	"github.com/Roh-Bot/Backend/models/Registration"
	"github.com/Roh-Bot/Backend/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

const ConnectString = `host=localhost port=5432 user=postgres password=admin dbname=User sslmode=disable`

var Register = &Registration.Register

func RegistrationMiddleware(c echo.Context) error {
	pool := utils.PostgresConnectionPool()

	v := validator.New()

	var response Registration.RegistrationResponse

	if err := c.Bind(&Register); err != nil {
		fmt.Println(err)
		fmt.Println("HI")
		return c.NoContent(http.StatusBadRequest)
	}
	if err := v.Struct(Register); err != nil {
		response = Registration.RegistrationResponse{
			StatusCode: 400,
			Error: map[string]string{
				"code":    "400",
				"message": "Something went wrong",
			},
			Data: map[string]string{},
		}
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	pgCheckIfUserExists := `SELECT email FROM users WHERE email='$1'`
	var email string
	errScan := pool.QueryRow(context.Background(), pgCheckIfUserExists, Register.Email).Scan(email)
	if errScan != nil {
		fmt.Println(errScan)
		fmt.Println("Scanning Error")
	}
	fmt.Println(email)

	passwordHash := utils.SHA256(Register.Password)
	fmt.Println(passwordHash)

	if utils.CheckIfEmailExists(Register.Email) {
		return c.String(404, "Email already exists")
	} else {
		pgCallStatement := `CALL registration($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`
		data := []any{Register.Email, Register.Phone_no, passwordHash, Register.First_name,
			Register.Last_name, Register.Dob, Register.Address_line_1, Register.Address_line_2,
			Register.City_id, Register.State_id, Register.Pincode, Register.Referred_by,
			Register.Reference_code}

		_, errQuery := pool.Query(context.Background(), pgCallStatement, data...)
		if errQuery != nil {
			fmt.Println("Query unsuccessful ", errQuery)
		}
		fmt.Println("Email of Middleware")
		fmt.Println(Register.Email)

		fmt.Println("Email of GET")
		utils.GetEmailFromDB()
		response = Registration.RegistrationResponse{
			StatusCode: 200,
			Error:      map[string]string{},
			Data: map[string]string{
				"user": "01",
			},
		}
		utils.EmailMagicLink(Register.User_Id, Register.Email)
		return c.JSON(http.StatusOK, response)

	}
}
