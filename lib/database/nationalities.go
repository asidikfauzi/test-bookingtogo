package database

import (
	"errors"
	"test-bookingtogo/config"
	"test-bookingtogo/models"
)

func GetNationalities(offset, limit int) ([]models.Nationalities, int64, error) {
	var nationalities []models.Nationalities
	var totalCount int64

	if err := config.DB.Offset(offset).Limit(limit).Find(&nationalities).Error; err != nil {
		return nationalities, totalCount, err
	}

	if err := config.DB.Model(&models.Nationalities{}).Count(&totalCount).Error; err != nil {
		return nationalities, totalCount, err
	}
	return nationalities, totalCount, nil

}

func GetNationalityById(id int) (models.Nationalities, error) {
	var nationalities models.Nationalities

	if rows := config.DB.Find(&nationalities, id).RowsAffected; rows < 1 {
		err := errors.New("Nationality not found!")
		return nationalities, err
	}
	return nationalities, nil
}

func InsertNationality(postBody models.NationalityPost) (interface{}, error) {
	nationality := models.Nationalities{
		NationalityName: postBody.NationalityName,
		NationalityCode: postBody.NationalityCode,
	}
	err := config.DB.Create(&nationality).Error
	return nationality, err
}
