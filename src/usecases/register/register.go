package registerUseCase

import (
	"errors"
	"os"
	"pembiayaan/src/entities"
	"pembiayaan/src/helpers"
	"pembiayaan/src/helpers/dates"
)

type (
	ResultToken struct {
		Name            string `json:"Name"`
		Email           string `json:"Email"`
		Role            string `json:"Role"`
		Token           string `json:"Token"`
		ExpiredInMinute string `json:"ExpiredIn"`
	}

	ParamRegister struct {
		Name     string
		Email    string
		Password string
		Role     string
	}
)

func (i *RegisterUseCase) Register(p *ParamRegister) (*ResultToken, error) {
	// check exist email
	_, err := i.RegisterRepository.FindByEmail(p.Email)

	if err == nil {
		return nil, helpers.ErrorMessage("0101", errors.New("email already exist"))
	}

	//hash password
	hashPassword, err := helpers.HashPassword(p.Password)

	if err != nil {
		return nil, helpers.ErrorMessage("0004", err)
	}

	//create user
	p.Password = hashPassword

	//role only admin, borrower, and lender
	if p.Role != "admin" && p.Role != "borrower" && p.Role != "lender" {
		p.Role = "borrower"
	}

	user := &entities.User{
		Name:      p.Name,
		Email:     p.Email,
		Password:  p.Password,
		Role:      p.Role,
		Status:    1,
		Image:     nil,
		CreatedAt: dates.GetCurrentTime(),
		UpdatedAt: dates.GetCurrentTime(),
	}

	_, err = i.RegisterRepository.Create(user)

	if err != nil {
		return nil, helpers.ErrorMessage("0004", err)
	}

	//generate api key
	apiKey, _ := helpers.GenerateApiKey(p.Password)

	paramsGenerateJWT := helpers.ParamsGenerateJWT{
		ExpiredInMinute: 60,
		SecretKey:       os.Getenv("ACCESS_TOKEN_SECRET_KEY"),
		Email:           p.Email,
		Name:            p.Name,
		Role:            p.Role,
		ApiKey:          apiKey,
	}

	token, exp, err := helpers.GenerateJWT(&paramsGenerateJWT)
	if err != nil {
		return nil, helpers.ErrorMessage("0004", err)
	}

	expInMinutes := dates.ConvertUnixToDatetime(exp)

	result := ResultToken{
		Name:            p.Name,
		Email:           p.Email,
		Role:            p.Role,
		Token:           token,
		ExpiredInMinute: expInMinutes,
	}

	return &result, nil

}
