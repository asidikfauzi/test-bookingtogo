package database

import (
	"errors"
	"gorm.io/gorm"
	"test-bookingtogo/config"
	"test-bookingtogo/models"
	"time"
)

func GetFamilyLists(offset, limit int) ([]models.FamilyListsResponse, int64, error) {
	var familyLists []models.FamilyLists
	var totalCount int64

	if err := config.DB.Table("family_lists").
		Preload("Customers").
		Select("family_lists.fl_id, family_lists.cst_id, family_lists.fl_relation, family_lists.fl_name, family_lists.fl_dob, customers.cst_name").
		Joins("JOIN customers ON family_lists.cst_id = customers.cst_id").
		Order("fl_name ASC").
		Offset(offset).
		Limit(limit).
		Find(&familyLists).Error; err != nil {
		return nil, totalCount, err
	}

	if err := config.DB.Model(&familyLists).Count(&totalCount).Error; err != nil {
		return nil, totalCount, err
	}

	var response []models.FamilyListsResponse
	for _, family := range familyLists {
		flDOB := family.FlDOB.Format("02-01-2006")
		response = append(response, models.FamilyListsResponse{
			FlId:       family.FlId,
			CstId:      family.CstId,
			CstName:    family.Customers.CstName,
			FlRelation: family.FlRelation,
			FlName:     family.FlName,
			FlDOB:      flDOB,
		})
	}

	return response, totalCount, nil

}

func GetFamilyListById(id int) (models.FamilyLists, error) {
	var familyLists models.FamilyLists

	if err := config.DB.Preload("Customers.Nationalities").First(&familyLists, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return familyLists, errors.New("FamilyLists not found")
		}
		return familyLists, err
	}
	return familyLists, nil
}

func InsertFamilyList(postBody models.FamilyListPost) (interface{}, error) {
	var familyLists models.FamilyLists

	flDOB, err := time.Parse("2006-01-02", postBody.FlDOB)
	if err != nil {
		e := errors.New("failed to parse data")
		return familyLists, e
	}

	familyLists.CstId = postBody.CstId
	familyLists.FlRelation = postBody.FlRelation
	familyLists.FlDOB = flDOB
	familyLists.FlName = postBody.FlName

	errResult := config.DB.Create(&familyLists).Error
	return postBody, errResult
}

func UpdateFamilyList(id int, putBody models.FamilyListUpdate) (interface{}, error) {
	var familyLists models.FamilyLists

	flDOB, err := time.Parse("2006-01-02", putBody.FlDOB)
	if err != nil {
		e := errors.New("failed to parse data")
		return familyLists, e
	}

	familyLists.CstId = putBody.CstId
	familyLists.FlRelation = putBody.FlRelation
	familyLists.FlDOB = flDOB
	familyLists.FlName = putBody.FlName

	errResult := config.DB.Model(&familyLists).Where("fl_id=?", id).Updates(&familyLists).Error
	return putBody, errResult
}

func DeleteFamilyList(id int) error {
	var familyLists models.FamilyLists
	result := config.DB.Delete(&familyLists, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
