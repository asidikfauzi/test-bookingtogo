package models

type Nationalities struct {
	NationalityId   int    `gorm:"primaryKey;AUTO_INCREMENT;" json:"nationality_id"`
	NationalityName string `gorm:"type:varchar(50);not null" json:"nationality_name"`
	NationalityCode string `gorm:"type:char(2);not null;" json:"nationality_code"`
}

type NationalityResponse struct {
	NationalityId   int    `json:"nationality_id"`
	NationalityName string `json:"nationality_name"`
	NationalityCode string `json:"nationality_code"`
}

type NationalityPost struct {
	NationalityName string `json:"nationality_name" form:"nationality_name" validate:"required"`
	NationalityCode string `json:"nationality_code" form:"nationality_code" validate:"required"`
}

type NationalityUpdate struct {
	NationalityName string `json:"nationality_name" form:"nationality_name" validate:"required"`
	NationalityCode string `json:"nationality_code" form:"nationality_code" validate:"required"`
}
