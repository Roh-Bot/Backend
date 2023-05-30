package Registration

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Verification(context echo.Context) error {
	return context.String(http.StatusOK, "Default Page")
}
