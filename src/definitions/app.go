package definitions

import (
	"gorm.io/gorm"
)

type (
	AppContext struct {
		Gorm *gorm.DB
	}
)
