package database

import (
	"errors"
	"test-bookingtogo/config"
	"test-bookingtogo/models"
)

func GetNationalities(offset, limit int) ([]models.NationalityResponse, int64, error) {
	var nationalities []models.Nationalities
	var totalCount int64

	if err := config.DB.Order("nationality_name ASC").Offset(offset).Limit(limit).Find(&nationalities).Error; err != nil {
		return nil, totalCount, err
	}

	if err := config.DB.Model(&nationalities).Count(&totalCount).Error; err != nil {
		return nil, totalCount, err
	}

	var response []models.NationalityResponse
	for _, nationality := range nationalities {
		response = append(response, models.NationalityResponse{
			NationalityId:   nationality.NationalityId,
			NationalityName: nationality.NationalityName,
			NationalityCode: nationality.NationalityCode,
		})
	}
	return response, totalCount, nil

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
	var nationality models.Nationalities

	nationality.NationalityName = postBody.NationalityName
	nationality.NationalityCode = postBody.NationalityCode

	err := config.DB.Create(&nationality).Error
	return nationality, err
}

func UpdateNationality(id int, putBody models.NationalityUpdate) (interface{}, error) {
	var nationality models.Nationalities

	nationality.NationalityName = putBody.NationalityName
	nationality.NationalityCode = putBody.NationalityCode

	err := config.DB.Model(&nationality).Where("nationality_id=?", id).Updates(&nationality).Error
	return putBody, err
}

func DeleteNationality(id int) error {
	var nationality models.Nationalities
	result := config.DB.Delete(&nationality, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
