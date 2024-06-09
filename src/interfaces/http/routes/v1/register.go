package v1routes

import RegisterV1Controller "pembiayaan/src/interfaces/http/controllers/v1/register"

func (i *V1Route) MountRegister() {
	g := i.Echo.Group("/auth")
	registerController := RegisterV1Controller.NewV1RegisterController(i.AppContext)

	g.POST("/register", registerController.Register)
}
