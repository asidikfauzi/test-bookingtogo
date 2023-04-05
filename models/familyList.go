package models

type FamilyLists struct {
	ID         int    `gorm:"primaryKey;AUTO_INCREMENT;" json:"fl_id"`
	CustomerId int    `gorm:"not null" json:"cst_id"`
	Relation   string `gorm:"type:varchar(50);not null;" json:"fl_relation"`
	Name       string `gorm:"type:varchar(50);not null;" json:"fl_name"`
	DOB        string `gorm:"type:varchar(50);not null;" json:"fl_name"`
}
