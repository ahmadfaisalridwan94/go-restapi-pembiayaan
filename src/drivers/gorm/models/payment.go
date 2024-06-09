package models

/*
CREATE TABLE IF NOT EXISTS payments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    loan_id INT NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    payment_date TIMESTAMP NOT NULL,
    status VARCHAR(50) NOT NULL, -- paid/pending/failed
    proof_of_payment_image VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (loan_id) REFERENCES loans(id)
);
*/

type Payment struct {
	Id                  int    `gorm:"primaryKey;column:id;type:int;not null;autoIncrement"`
	LoanId              int    `gorm:"column:loan_id;type:int;not null"`
	Amount              string `gorm:"column:amount;type:decimal(10,2);not null"`
	PaymentDate         string `gorm:"column:payment_date;type:timestamp;not null"`
	Status              string `gorm:"column:status;type:varchar(50);not null"`
	ProofOfPaymentImage string `gorm:"column:proof_of_payment_image;type:varchar(255)"`
	CreatedAt           string `gorm:"column:created_at;type:timestamp;not null"`
}

func (Payment) TableName() string {
	return "payments"
}
