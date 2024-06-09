//go:build wireinject
// +build wireinject

package registerUseCase

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeRegisterUseCase(
	gormDb *gorm.DB,
	ctx context.Context,
) IRegisterUseCase {
	wire.Build(
		NewRegisterUseCase,
		NewRegisterRepository,
	)

	return nil
}
