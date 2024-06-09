//go:build wireinject
// +build wireinject

package loginUseCase

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeLoginUseCase(
	gormDb *gorm.DB,
	ctx context.Context,
) ILoginUseCase {
	wire.Build(
		NewLoginUseCase,
		NewLoginRepository,
	)

	return nil
}
