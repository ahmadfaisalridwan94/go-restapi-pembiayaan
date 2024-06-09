// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package loginUseCase

import (
	"context"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeLoginUseCase(gormDb *gorm.DB, ctx context.Context) ILoginUseCase {
	iLoginRepository := NewLoginRepository(gormDb, ctx)
	iLoginUseCase := NewLoginUseCase(iLoginRepository, gormDb, ctx)
	return iLoginUseCase
}