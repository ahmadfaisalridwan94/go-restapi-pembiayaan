package registerRepository

import (
	"context"
	"pembiayaan/src/entities"

	"gorm.io/gorm"
)

type RegisterRepository struct {
	Gorm    *gorm.DB
	Context context.Context
}

var CollectionName = "Register"

type IRegisterRepository interface {
	Create(param *entities.User) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
}

func NewRegisterRepository(gormDB *gorm.DB, ctx context.Context) IRegisterRepository {
	return &RegisterRepository{
		Gorm:    gormDB,
		Context: ctx,
	}
}
