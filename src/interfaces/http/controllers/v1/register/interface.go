package registerV1Controller

import (
	"pembiayaan/src/definitions"

	"github.com/labstack/echo/v4"
)

type V1RegisterController struct {
	*definitions.AppContext
}

type IV1RegisterController interface {
	Register(c echo.Context) error
}

func NewV1RegisterController(appContext *definitions.AppContext) IV1RegisterController {
	return &V1RegisterController{
		appContext,
	}
}
