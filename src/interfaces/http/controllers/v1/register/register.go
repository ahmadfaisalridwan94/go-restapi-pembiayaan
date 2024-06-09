package registerV1Controller

import (
	"net/http"
	"pembiayaan/src/definitions"
	"pembiayaan/src/helpers"
	registerUC "pembiayaan/src/usecases/register"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type (
	reqRegister struct {
		Name     string `json:"Name" validate:"required"`
		Email    string `json:"Email" validate:"required"`
		Password string `json:"Password" validate:"required"`
		Role     string `json:"Role" validate:"required"`
	}
)

func (i *V1RegisterController) Register(c echo.Context) (err error) {

	payload := reqRegister{}
	if err = c.Bind(&payload); err != nil {
		return helpers.ErrorMessage("0013", err)
	}

	if err = c.Validate(&payload); err != nil {
		return helpers.ErrorMessage("0013", err)
	}

	registerUseCase := registerUC.InitializeRegisterUseCase(i.Gorm, c.Request().Context())

	paramRegister := registerUC.ParamRegister{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
		Role:     payload.Role,
	}

	result, err := registerUseCase.Register(&paramRegister)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, definitions.SuccessResponse{
		ResponseCode: "0000",
		ResponseDesc: "Success",
		ResponseData: result,
	})
}

// Validate validates the request payload using the validator library
func (r *reqRegister) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
