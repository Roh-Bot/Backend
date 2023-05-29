package handlers

import (
	"github.com/Roh-Bot/Backend/middlewares"
	"github.com/labstack/echo/v4"
	"log"
)

func Start() {
	router := echo.New()
	router.GET("/", middlewares.DefaultPageMiddleware)
	router.POST("strategy/dss", middlewares.DSSMiddleware)
	router.POST("/register", middlewares.RegistrationController)
	router.POST("/login", middlewares.LoginController)
	router.GET("/home", middlewares.Home)
	router.POST("/loginAuth", middlewares.Login)
	router.POST("/refresh", middlewares.Refresh)
	log.Fatal(router.Start("localhost:8080"))
}
