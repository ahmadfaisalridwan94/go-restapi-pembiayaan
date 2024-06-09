package registerUseCase

import (
	"context"
	registerRepository "pembiayaan/src/repositories/register"

	"gorm.io/gorm"
)

type (
	RegisterUseCase struct {
		RegisterRepository registerRepository.IRegisterRepository
		Gorm               *gorm.DB
		Context            context.Context
	}

	IRegisterUseCase interface {
		Register(p *ParamRegister) (*ResultToken, error)
	}
)

func NewRegisterRepository(
	gormDb *gorm.DB,
	ctx context.Context,
) registerRepository.IRegisterRepository {
	return registerRepository.NewRegisterRepository(gormDb, ctx)
}

func NewRegisterUseCase(
	registerRepository registerRepository.IRegisterRepository,
	gorm *gorm.DB,
	ctx context.Context,
) IRegisterUseCase {
	return &RegisterUseCase{
		RegisterRepository: registerRepository,
		Gorm:               gorm,
		Context:            ctx,
	}
}
