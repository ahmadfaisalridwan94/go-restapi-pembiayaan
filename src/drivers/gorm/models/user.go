package models

import "time"

/*
users (
	id serial PRIMARY KEY,
	name varchar(255) NOT NULL,
	email varchar(255) NOT NULL UNIQUE,
	email_verify_at timestamp NULL,
	password varchar(255) NOT NULL,
	remember_token varchar(100) NULL,
	image varchar(255) NULL,
	role varchar(255) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL
);
*/
type User struct {
	Id            int        `gorm:"primaryKey;column:id;type:int;not null;autoIncrement"`
	Name          string     `gorm:"column:name;type:varchar(255);not null"`
	Email         string     `gorm:"column:email;type:varchar(255);not null"`
	EmailVerifyAt *time.Time `gorm:"column:email_verify_at;type:timestamp"`
	Password      string     `gorm:"column:password;type:varchar(255);not null"`
	RememberToken *string    `gorm:"column:remember_token;type:varchar(100)"`
	Image         *string    `gorm:"column:image;type:varchar(255)"`
	Role          string     `gorm:"column:role;type:varchar(255);not null"`
	Status        int        `gorm:"column:status;type:varchar(50);not null"`
	CreatedAt     time.Time  `gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt     time.Time  `gorm:"column:updated_at;type:timestamp;not null"`
}

func (User) TableName() string {
	return "users"
}
