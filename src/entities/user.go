package entities

import (
	"pembiayaan/src/drivers/gorm/models"
	"time"
)

type User struct {
	Id        int
	Name      string
	Email     string
	Image     *string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToUserEntity(user *models.User) *User {
	return &User{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Image:     user.Image,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
