package entities

import (
	"pembiayaan/src/drivers/gorm/models"
	"time"
)

type User struct {
	Id        int
	Name      string
	Password  string
	Email     string
	Image     *string
	Role      string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToUserEntity(user *models.User) *User {
	return &User{
		Id:        user.Id,
		Name:      user.Name,
		Password:  user.Password,
		Email:     user.Email,
		Image:     user.Image,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
