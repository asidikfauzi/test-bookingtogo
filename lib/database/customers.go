package database

import (
	"errors"
	"test-bookingtogo/config"
	"test-bookingtogo/models"
	"time"
)

func GetCustomers(offset, limit int) ([]models.Customers, int64, error) {
	var customers []models.Customers
	var totalCount int64

	if err := config.DB.Preload("Nationalities").Order("cst_name ASC").Offset(offset).Limit(limit).Find(&customers).Error; err != nil {
		return customers, totalCount, err
	}

	if err := config.DB.Model(&customers).Count(&totalCount).Error; err != nil {
		return customers, totalCount, err
	}
	return customers, totalCount, nil

}

func GetCustomerById(id int) (models.Customers, error) {
	var customers models.Customers

	if rows := config.DB.Preload("Nationalities").Find(&customers, id).RowsAffected; rows < 1 {
		err := errors.New("Customer not found!")
		return customers, err
	}
	return customers, nil
}

func InsertCustomer(postBody models.CustomerPost) (interface{}, error) {
	var customers models.Customers

	cstDOB, err := time.Parse("02/01/2006", postBody.CstDOB)
	if err != nil {
		e := errors.New("Failed to parse data!")
		return customers, e
	}

	customers.NationalityId = postBody.NationalityId
	customers.CstName = postBody.CstName
	customers.CstDOB = cstDOB
	customers.CstPhoneNum = postBody.CstPhoneNum
	customers.CstEmail = postBody.CstEmail

	errResult := config.DB.Create(&customers).Error
	return postBody, errResult
}

func UpdateCustomer(id int, putBody models.CustomerUpdate) (interface{}, error) {
	var customers models.Customers

	cstDOB, err := time.Parse("02/01/2006", putBody.CstDOB)
	if err != nil {
		e := errors.New("Failed to parse data!")
		return customers, e
	}

	customers.NationalityId = putBody.NationalityId
	customers.CstName = putBody.CstName
	customers.CstDOB = cstDOB
	customers.CstPhoneNum = putBody.CstPhoneNum
	customers.CstEmail = putBody.CstEmail

	errResult := config.DB.Model(&customers).Where("cst_id=?", id).Updates(&customers).Error
	return putBody, errResult
}

func DeleteCustomer(id int) error {
	var customers models.Customers
	result := config.DB.Delete(&customers, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
