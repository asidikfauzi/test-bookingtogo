package models

type FamilyLists struct {
	FlId       int       `gorm:"primaryKey;AUTO_INCREMENT;" json:"fl_id"`
	CustomerId int       `gorm:"not null" json:"cst_id"`
	FlRelation string    `gorm:"type:varchar(50);not null;" json:"fl_relation"`
	FlName     string    `gorm:"type:varchar(50);not null;" json:"fl_name"`
	FlDOB      string    `gorm:"type:varchar(50);not null;" json:"fl_name"`
	Customers  Customers `gorm:"foreignKey:CstId;references:CustomerId"`
}
