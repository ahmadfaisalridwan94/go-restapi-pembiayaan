package entities

import (
	"pembiayaan/src/drivers/gorm/models"
)

type Payment struct {
	Id                  int
	LoanId              int
	Amount              string
	PaymentDate         string
	Status              string
	ProofOfPaymentImage string
	CreatedAt           string
}

func ToPaymentEntity(payment *models.Payment) *Payment {
	return &Payment{
		Id:                  payment.Id,
		LoanId:              payment.LoanId,
		Amount:              payment.Amount,
		PaymentDate:         payment.PaymentDate,
		Status:              payment.Status,
		ProofOfPaymentImage: payment.ProofOfPaymentImage,
		CreatedAt:           payment.CreatedAt,
	}
}
