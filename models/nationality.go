package models

type Nationality struct {
	ID   int    `gorm:"primaryKey;AUTO_INCREMENT;" json:"nationality_id"`
	Name string `gorm:"type:varchar(50);not null" json:"nationality_name"`
	Code string `gorm:"type:char(2);not null;" json:"nationality_code"`
}
