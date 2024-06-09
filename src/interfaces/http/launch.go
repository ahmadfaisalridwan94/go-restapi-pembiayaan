package http

import (
	"pembiayaan/src/helpers"
	"pembiayaan/src/helpers/validators"
	v1routes "pembiayaan/src/interfaces/http/routes/v1"

	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (i *Http) Launch() {
	e := echo.New()

	var corsConfigs []string

	if os.Getenv("CORS") != "" {
		corsConfigs = strings.Split(os.Getenv("CORS"), ",")
	}

	e.Validator = &validators.CustomValidator{Validator: validators.InitValidator()}
	e.HTTPErrorHandler = helpers.ErrorHandler
	e.Pre(middleware.RemoveTrailingSlash())
	if len(corsConfigs) > 0 {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: corsConfigs,
		}))
	}

	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.RequestID())
	e.Use(middleware.BodyLimit("3M"))

	basePath := e.Group(os.Getenv("BASE_PATH"))

	basePath.GET("", func(c echo.Context) error {
		return c.HTML(http.StatusOK, fmt.Sprintf("API User for %s %s", os.Getenv("ENVIRONMENT"), os.Getenv("CORS")))
	})

	v1 := v1routes.NewV1Route(i.AppContext, basePath.Group("/v1"))
	// v2 := v2routes.NewV2Route(i.AppContext, basePath.Group("/v2"))
	basePath.GET("/api-docs/*", echoSwagger.WrapHandler)

	// routes
	v1.MountPing()
	v1.MountLogin()

	// v2.MountPing()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${id} | ${remote_ip} | ${host} | ${method} | ${uri} | ${user_agent} | ${status} | ${error} | ${latency_human} | ${bytes_in} | ${bytes_out} |\n",
		Output: os.Stdout,
	}))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))))
}
