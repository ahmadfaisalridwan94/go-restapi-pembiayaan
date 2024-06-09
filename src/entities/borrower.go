package entities

import (
	"pembiayaan/src/drivers/gorm/models"
	"time"
)

type Borrower struct {
	UserId      int
	Address     string
	DateOfBirth time.Time
	IdNumber    string
	TaxNumber   string
}

func ToBorrowerEntity(borrower *models.Borrower) *Borrower {
	return &Borrower{
		UserId:      borrower.UserId,
		Address:     borrower.Address,
		DateOfBirth: borrower.DateOfBirth,
		IdNumber:    borrower.IdNumber,
		TaxNumber:   borrower.TaxNumber,
	}
}
