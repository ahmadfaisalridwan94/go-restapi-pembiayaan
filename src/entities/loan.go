package entities

import (
	"pembiayaan/src/drivers/gorm/models"
	"time"
)

type Loan struct {
	Id         int
	BorrowerId int
	LenderId   int
	Amount     string
	Interest   string
	Term       int
	Status     string
	CreatedAt  time.Time
	ApprovedAt time.Time
}

func ToLoanEntity(loan *models.Loan) *Loan {
	return &Loan{
		Id:         loan.Id,
		BorrowerId: loan.BorrowerId,
		LenderId:   loan.LenderId,
		Amount:     loan.Amount,
		Interest:   loan.Interest,
		Term:       loan.Term,
		Status:     loan.Status,
		CreatedAt:  loan.CreatedAt,
		ApprovedAt: loan.ApprovedAt,
	}
}
