package models

/*
CREATE TABLE IF NOT EXISTS lenders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    company_name VARCHAR(100),
    address TEXT,
    license_number VARCHAR(50),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
*/

type Lender struct {
	Id            int    `gorm:"primaryKey;column:id;type:int;not null;autoIncrement"`
	UserId        int    `gorm:"column:user_id;type:int;not null"`
	CompanyName   string `gorm:"column:company_name;type:varchar(100)"`
	Address       string `gorm:"column:address;type:text"`
	LicenseNumber string `gorm:"column:license_number;type:varchar(50)"`
}

func (Lender) TableName() string {
	return "lenders"
}
