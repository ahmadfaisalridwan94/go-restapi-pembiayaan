package loginRepository

import (
	"pembiayaan/src/entities"
)

func (i *LoginRepository) FindByEmailAndStatus(email string, status int) (*entities.User, error) {

	var result *entities.User
	err := i.Gorm.Table("users").Where("email = ? AND status = ?", email, status).First(&result).Error

	return result, err

}
