package database

import (
	"test-bookingtogo/config"
	"test-bookingtogo/models"
)

func InsertNationality(postBody models.NationalityPost) (interface{}, error) {
	nationality := models.Nationalities{
		NationalityName: postBody.NationalityName,
		NationalityCode: postBody.NationalityCode,
	}
	err := config.DB.Create(&nationality).Error
	return nationality, err
}
