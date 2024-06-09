package models

import "time"

/*
CREATE TABLE IF NOT EXISTS borrowers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    address TEXT NOT NULL,
    date_of_birth DATE NOT NULL,
    id_number VARCHAR(50) NOT NULL,
    tax_number VARCHAR(50),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
*/
type Borrower struct {
	Id          int       `gorm:"primaryKey;column:id;type:int;not null;autoIncrement"`
	UserId      int       `gorm:"column:user_id;type:int;not null"`
	Address     string    `gorm:"column:address;type:text;not null"`
	DateOfBirth time.Time `gorm:"column:date_of_birth;type:date;not null"`
	IdNumber    string    `gorm:"column:id_number;type:varchar(50);not null"`
	TaxNumber   string    `gorm:"column:tax_number;type:varchar(50)"`
}

func (Borrower) TableName() string {
	return "borrowers"
}
