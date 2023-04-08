package models

import "time"

type FamilyLists struct {
	FlId       int       `gorm:"primaryKey;AUTO_INCREMENT;" json:"fl_id"`
	CstId      int       `gorm:"not null" json:"cst_id"`
	FlRelation string    `gorm:"type:varchar(50);not null;" json:"fl_relation"`
	FlName     string    `gorm:"type:varchar(50);not null;" json:"fl_name"`
	FlDOB      time.Time `gorm:"not null;" json:"fl_dob"`
	Customers  Customers `gorm:"foreignKey:CstId;references:CstId" json:"customer"`
}

type FamilyListsResponse struct {
	FlId       int    `json:"fl_id"`
	CstId      int    `json:"cst_id"`
	CstName    string `json:"cst_name"`
	FlRelation string `json:"fl_relation"`
	FlName     string `json:"fl_name"`
	FlDOB      string `json:"fl_dob"`
}

type FamilyListPost struct {
	CstId      int    `json:"cst_id" form:"cst_id" validate:"required"`
	FlRelation string `json:"fl_relation" form:"fl_relation" validate:"required"`
	FlName     string `json:"fl_name" form:"fl_name" validate:"required"`
	FlDOB      string `json:"fl_dob" form:"fl_dob" validate:"required"`
}

type FamilyListUpdate struct {
	CstId      int    `json:"cst_id" form:"cst_id" validate:"required"`
	FlRelation string `json:"fl_relation" form:"fl_relation" validate:"required"`
	FlName     string `json:"fl_name" form:"fl_name" validate:"required"`
	FlDOB      string `json:"fl_dob" form:"fl_dob" validate:"required"`
}
