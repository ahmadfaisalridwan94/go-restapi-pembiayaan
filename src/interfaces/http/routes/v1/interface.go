package v1routes

import (
	"pembiayaan/src/definitions"

	"github.com/labstack/echo/v4"
)

type V1Route struct {
	*definitions.AppContext
	Echo *echo.Group
}

type IV1Route interface {
	MountPing()
	MountLogin()
}

func NewV1Route(appContext *definitions.AppContext, echoGroup *echo.Group) IV1Route {
	return &V1Route{
		appContext,
		echoGroup,
	}
}
