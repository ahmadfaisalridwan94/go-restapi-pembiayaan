package loginRepository

import (
	"context"
	"pembiayaan/src/entities"

	"gorm.io/gorm"
)

type LoginRepository struct {
	Gorm    *gorm.DB
	Context context.Context
}

var CollectionName = "Login"

type ILoginRepository interface {
	FindByEmailAndStatus(email string, status int) (*entities.User, error)
}

func NewLoginRepository(gormDB *gorm.DB, ctx context.Context) ILoginRepository {
	return &LoginRepository{
		Gorm:    gormDB,
		Context: ctx,
	}
}
