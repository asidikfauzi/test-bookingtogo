package models

import "time"

type Customer struct {
	ID        int        `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Username  string     `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email     string     `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string     `gorm:"type:varchar(255);not null" json:"password"`
	Address   string     `gorm:"type:varchar(255);" json:"address"`
	CreatedAt time.Time  `gorm:"type:varchar(255);" json:"address"`
	UpdatedAt time.Time  `gorm:"type:varchar(255);" json:"address"`
	DeletedAt *time.Time `gorm:"type:varchar(255);" json:"address"`
}
