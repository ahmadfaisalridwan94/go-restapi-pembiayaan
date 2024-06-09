package loginUseCase

import (
	"errors"
	"os"
	"pembiayaan/src/helpers"
	"pembiayaan/src/helpers/dates"
	"strconv"
)

type (
	ResultToken struct {
		Name            string `json:"Name"`
		Email           string `json:"Email"`
		Role            string `json:"Role"`
		Token           string `json:"Token"`
		ExpiredInMinute string `json:"ExpiredIn"`
	}

	ParamLogin struct {
		Email    string
		Password string
	}
)

func (i *LoginUseCase) Login(p *ParamLogin) (*ResultToken, error) {

	result, err := i.LoginRepository.FindByEmailAndStatus(p.Email, 1)
	if err != nil {
		return nil, helpers.ErrorMessage("0105", errors.New("Wrong credentials"))
	}

	//check status client
	if result.Status == 0 {
		return nil, helpers.ErrorMessage("0106", errors.New("Inactive Account"))
	}

	//validate password
	checkHash := helpers.CheckPasswordHash(p.Password, result.Password)
	if checkHash == false {
		return nil, helpers.ErrorMessage("0105", errors.New("Wrong credentials"))
	}

	envExp := os.Getenv("TOKEN_EXPIRED_IN_MINUTES")

	expiredToken, err := strconv.Atoi(envExp)
	if err != nil {
		return nil, helpers.ErrorMessage("0004", errors.New("Unauthorized"))
	}

	//generate api key
	apiKey, _ := helpers.GenerateApiKey(p.Password)

	paramsGenerateJWT := helpers.ParamsGenerateJWT{
		ExpiredInMinute: expiredToken,
		SecretKey:       os.Getenv("ACCESS_TOKEN_SECRET_KEY"),
		Email:           result.Email,
		Name:            result.Name,
		Role:            result.Role,
		ApiKey:          apiKey,
	}

	token, exp, err := helpers.GenerateJWT(&paramsGenerateJWT)
	if err != nil {
		return nil, helpers.ErrorMessage("0004", errors.New("Unauthorized"))
	}

	expInMinutes := dates.ConvertUnixToDatetime(exp)
	return &ResultToken{
		Name:            result.Name,
		Email:           result.Email,
		Role:            result.Role,
		Token:           token,
		ExpiredInMinute: expInMinutes,
	}, nil
}
