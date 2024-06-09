package loginV1Controller

import (
	"fmt"
	"net/http"
	"pembiayaan/src/definitions"
	"pembiayaan/src/helpers"
	loginUC "pembiayaan/src/usecases/login"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type (
	reqLogin struct {
		Email    string `json:"Email" validate:"required"`
		Password string `json:"Password" validate:"required"`
	}
)

func (i *V1LoginController) Login(c echo.Context) (err error) {

	fmt.Println("Login")

	payload := reqLogin{}
	if err = c.Bind(&payload); err != nil {
		return helpers.ErrorMessage("0013", err)
	}

	if err = c.Validate(&payload); err != nil {
		return helpers.ErrorMessage("0013", err)
	}

	loginUseCase := loginUC.InitializeLoginUseCase(i.Gorm, c.Request().Context())

	paramLogin := loginUC.ParamLogin{
		Email:    payload.Email,
		Password: payload.Password,
	}

	result, err := loginUseCase.Login(&paramLogin)
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
func (r *reqLogin) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
