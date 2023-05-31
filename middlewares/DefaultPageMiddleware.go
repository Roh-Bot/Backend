package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func DefaultPageMiddleware(context echo.Context) error {
	return context.String(http.StatusOK, "Registered Successfully")
}
