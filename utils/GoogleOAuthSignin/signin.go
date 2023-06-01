package GoogleOAuthSignin

import (
	"encoding/json"
	"fmt"
	"github.com/Roh-Bot/Backend/models/GoogleOAuth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var oauthGoogle *oauth2.Config
var RandomString = "random-string-" //random string can be generated using generater function
var e = echo.New()
var userInfo GoogleOAuth.UserInfoStruct

// func for initializing oauth
func init() {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	os.Setenv("GoogleKey", "818962244597-jirbmufc1lbhmovm6t4vo7h4tv3i3iut.apps.googleusercontent.com")
	os.Setenv("GoogleSecretKey", "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")

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	oauthGoogle = &oauth2.Config{
		RedirectURL:  "http://localhost:8000/callback",
		ClientID:     os.Getenv("GoogleKey"),
		ClientSecret: os.Getenv("GoogleSecretKey"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func Signin(context echo.Context) error {
	url := oauthGoogle.AuthCodeURL(RandomString)
	//fmt.Println(url)
	return context.Redirect(http.StatusTemporaryRedirect, url)
}

func HandleGoogleCallback(context echo.Context) error {
	content, err := getUserInfo(context.FormValue("state"), context.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		return context.Redirect(http.StatusTemporaryRedirect, "http://localhost:8000")
	}
	fmt.Println(content)
	return context.JSON(200, userInfo)
}

func getUserInfo(state string, code string) ([]byte, error) {
	token, err := oauthGoogle.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	fmt.Println(token.AccessToken)
	response, err := http.Get("https://www.googleapis.com/userinfo/v2/me?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	var mydata2 = string(body)
	fmt.Println("Mydata2", mydata2)

	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, fmt.Errorf("failed parsing response body: %s", err.Error())
	}
	fmt.Println(userInfo)
	return nil, nil
}
