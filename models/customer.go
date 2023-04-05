package models

import "time"

type Customer struct {
	ID            int         `gorm:"primaryKey;AUTO_INCREMENT" json:"cst_id"`
	NationalityId int         `gorm:"not null" json:"nationality_id"`
	DOB           time.Time   `gorm:"not null;" json:"cst_dob"`
	PhoneNum      string      `gorm:"type:varchar(20);not null" json:"cst_phoneNum"`
	Email         string      `gorm:"type:varchar(50);" json:"cst_email"`
	Nationality   Nationality `gorm:"foreignKey:ID;references:NationalityId"`
}
