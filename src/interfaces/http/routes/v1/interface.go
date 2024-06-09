package v1routes

import (
	"pembiayaan/src/definitions"
)

type V1Route struct {
	*definitions.AppContext
	Echo *echo.group
}
