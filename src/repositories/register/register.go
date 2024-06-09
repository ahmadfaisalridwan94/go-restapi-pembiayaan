package registerRepository

import (
	"pembiayaan/src/entities"
)

// function for create users
func (i *RegisterRepository) Create(param *entities.User) (*entities.User, error) {

	//using table name for creating user data
	err := i.Gorm.Table("users").Create(&param).Error
	if err != nil {
		return nil, err
	}

	return param, nil
}

// function find by email
func (i *RegisterRepository) FindByEmail(email string) (*entities.User, error) {

	var result *entities.User
	err := i.Gorm.Where("email = ?", email).First(&result).Error

	return result, err
}
