package v1routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *V1Route) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}
