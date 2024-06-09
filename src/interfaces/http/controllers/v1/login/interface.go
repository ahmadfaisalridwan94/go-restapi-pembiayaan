package loginV1Controller

import (
	"pembiayaan/src/definitions"

	"github.com/labstack/echo/v4"
)

type V1LoginController struct {
	*definitions.AppContext
}

type IV1LoginController interface {
	Login(c echo.Context) error
}

func NewV1LoginController(appContext *definitions.AppContext) IV1LoginController {
	return &V1LoginController{
		appContext,
	}
}
