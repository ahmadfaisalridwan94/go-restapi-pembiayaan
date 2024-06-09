package v1controller

import (
	"pembiayaan/src/definitions"

	"github.com/labstack/echo/v4"
)

type V1Controller struct {
	*definitions.AppContext
}

type IV1Controller interface {
	Ping(c echo.Context) error
}

func NewV1Controller(appContext *definitions.AppContext) IV1Controller {
	return &V1Controller{
		appContext,
	}
}
