package models

import "time"

type Customers struct {
	CstId         int           `gorm:"primaryKey;AUTO_INCREMENT" json:"cst_id"`
	NationalityId int           `gorm:"not null" json:"nationality_id"`
	CstName       string        `gorm:"type:char(50);not null" json:"cst_name"`
	CstDOB        time.Time     `gorm:"not null;" json:"cst_dob"`
	CstPhoneNum   string        `gorm:"type:varchar(20);not null" json:"cst_phone_num"`
	CstEmail      string        `gorm:"type:varchar(50);" json:"cst_email"`
	Nationalities Nationalities `gorm:"foreignKey:NationalityId;references:NationalityId;" json:"nationality"`
}

type CustomersResponse struct {
	CstId           int    `json:"cst_id"`
	NationalityId   int    `json:"nationality_id"`
	NationalityName string `json:"nationality_name"`
	CstName         string `json:"cst_name"`
	CstDOB          string `json:"cst_dob"`
	CstPhoneNum     string `json:"cst_phone_num"`
	CstEmail        string `json:"cst_email"`
}

type CustomerPost struct {
	NationalityId int    `json:"nationality_id" form:"nationality_id" validate:"required"`
	CstName       string `json:"cst_name" form:"cst_name" validate:"required"`
	CstDOB        string `json:"cst_dob" form:"cst_dob" validate:"required"`
	CstPhoneNum   string `json:"cst_phone_num" form:"cst_phone_num" validate:"required"`
	CstEmail      string `json:"cst_email" form:"cst_email" validate:"required"`
}

type CustomerUpdate struct {
	NationalityId int    `json:"nationality_id" form:"nationality_id" validate:"required"`
	CstName       string `json:"cst_name" form:"cst_name" validate:"required"`
	CstDOB        string `json:"cst_dob" form:"cst_dob" validate:"required"`
	CstPhoneNum   string `json:"cst_phone_num" form:"cst_phone_num" validate:"required"`
	CstEmail      string `json:"cst_email" form:"cst_email" validate:"required"`
}
