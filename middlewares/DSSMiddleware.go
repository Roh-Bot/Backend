package middlewares

import (
	"github.com/Roh-Bot/Backend/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

func DSSMiddleware(context echo.Context) error {

	v := validator.New()
	var dss models.DSS

	if err := context.Bind(&dss); err != nil {
		return context.String(http.StatusInternalServerError, "Binding Error")
	}

	if err := v.Struct(dss); err != nil {
		return context.String(http.StatusInternalServerError, "Validation Error")
	}
	context.Response().Header().Add("session_key", "ASDT16461ASD")

	return context.JSON(http.StatusOK, dss)

}
