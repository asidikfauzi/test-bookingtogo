package models

import "time"

type Customers struct {
	CstId         int           `gorm:"primaryKey;AUTO_INCREMENT" json:"cst_id"`
	NationalityId int           `gorm:"not null" json:"nationality_id"`
	CstDOB        time.Time     `gorm:"not null;" json:"cst_dob"`
	CstPhoneNum   string        `gorm:"type:varchar(20);not null" json:"cst_phoneNum"`
	CstEmail      string        `gorm:"type:varchar(50);" json:"cst_email"`
	Nationalities Nationalities `gorm:"foreignKey:NationalityId;references:NationalityId"`
}
