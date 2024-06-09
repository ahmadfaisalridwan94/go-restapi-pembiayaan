package loginUseCase

import (
	"context"
	loginRepository "pembiayaan/src/repositories/login"

	"gorm.io/gorm"
)

type (
	LoginUseCase struct {
		LoginRepository loginRepository.ILoginRepository
		Gorm            *gorm.DB
		Context         context.Context
	}

	ILoginUseCase interface {
		Login(p *ParamLogin) (*ResultToken, error)
	}
)

func NewLoginRepository(
	gormDb *gorm.DB,
	ctx context.Context,
) loginRepository.ILoginRepository {
	return loginRepository.NewLoginRepository(gormDb, ctx)
}

func NewLoginUseCase(
	loginRepository loginRepository.ILoginRepository,
	gorm *gorm.DB,
	ctx context.Context,
) ILoginUseCase {
	return &LoginUseCase{
		LoginRepository: loginRepository,
		Gorm:            gorm,
		Context:         ctx,
	}
}
