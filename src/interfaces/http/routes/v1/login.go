package v1routes

import LoginV1Controller "pembiayaan/src/interfaces/http/controllers/v1/login"

func (i *V1Route) MountLogin() {
	g := i.Echo.Group("/auth")
	loginController := LoginV1Controller.NewV1LoginController(i.AppContext)

	g.POST("/login", loginController.Login)
}
