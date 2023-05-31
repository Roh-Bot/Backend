package handlers

import (
	"github.com/Roh-Bot/Backend/middlewares"
	"github.com/Roh-Bot/Backend/middlewares/Authentication"
	"github.com/Roh-Bot/Backend/middlewares/DSS"
	"github.com/Roh-Bot/Backend/middlewares/GoogleOAuthSignin"
	"github.com/Roh-Bot/Backend/middlewares/Login"
	"github.com/Roh-Bot/Backend/middlewares/Registration"
	"github.com/Roh-Bot/Backend/middlewares/SessionHandling"
	"github.com/labstack/echo/v4"
	"log"
)

func Start() {
	router := echo.New()
	router.GET("/", middlewares.DefaultPageMiddleware)
	router.POST("strategy/dss", DSS.DSSMiddleware)
	router.POST("/register", Registration.RegistrationMiddleware)
	router.GET("/verifylink", Registration.Verification)
	router.POST("/login", Login.LoginController)
	router.GET("/home", Authentication.Home)
	router.POST("/loginAuth", Authentication.Login)
	router.POST("/refresh", Authentication.Refresh)
	router.POST("/loginSession", SessionHandling.Signin)
	router.GET("/signin", GoogleOAuthSignin.Signin)
	router.GET("/callback", GoogleOAuthSignin.HandleGoogleCallback)
	log.Fatal(router.Start("localhost:8000"))
}
