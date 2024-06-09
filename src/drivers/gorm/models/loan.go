package models

import "time"

/*
CREATE TABLE IF NOT EXISTS loans (
    id INT AUTO_INCREMENT PRIMARY KEY,
    borrower_id INT NOT NULL,
    lender_id INT NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    interest_rate DECIMAL(5,2) NOT NULL,
    term INT NOT NULL, -- in months
    status VARCHAR(50) NOT NULL, -- pending/approved/rejected
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    approved_at TIMESTAMP NULL,
    FOREIGN KEY (borrower_id) REFERENCES borrowers(id),
    FOREIGN KEY (lender_id) REFERENCES lenders(id)
);
*/

type Loan struct {
	Id         int       `gorm:"primaryKey;column:id;type:int;not null;autoIncrement"`
	BorrowerId int       `gorm:"column:borrower_id;type:int;not null"`
	LenderId   int       `gorm:"column:lender_id;type:int;not null"`
	Amount     string    `gorm:"column:amount;type:decimal(10,2);not null"`
	Interest   string    `gorm:"column:interest;type:decimal(5,2);not null"`
	Term       int       `gorm:"column:term;type:int;not null"`
	Status     string    `gorm:"column:status;type:varchar(50);not null"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp;not null"`
	ApprovedAt time.Time `gorm:"column:approved_at;type:timestamp"`
}
