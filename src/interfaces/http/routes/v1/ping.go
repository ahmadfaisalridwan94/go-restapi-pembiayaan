package v1routes

import (
	v1controller "pembiayaan/src/interfaces/http/controllers/v1"
)

func (i *V1Route) MountPing() {
	g := i.Echo.Group("/ping")
	pingController := v1controller.NewV1Controller(i.AppContext)

	g.GET("", pingController.Ping)
}
