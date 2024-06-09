package entities

import "pembiayaan/src/drivers/gorm/models"

type Lender struct {
	Id            int
	UserId        int
	CompanyName   string
	Address       string
	LicenseNumber string
}

func ToLenderEntity(l *models.Lender) *Lender {
	return &Lender{
		Id:            l.Id,
		UserId:        l.UserId,
		CompanyName:   l.CompanyName,
		Address:       l.Address,
		LicenseNumber: l.LicenseNumber,
	}
}
