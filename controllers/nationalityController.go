package controllers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
	"test-bookingtogo/config"
	"test-bookingtogo/lib/database"
	"test-bookingtogo/lib/utils"
	"test-bookingtogo/models"
)

var nationalities []models.Nationalities

func GetNationalities(w http.ResponseWriter, r *http.Request) {

	return
}

func CreateNationality(w http.ResponseWriter, r *http.Request) {
	var postBody models.NationalityPost

	if err := json.NewDecoder(r.Body).Decode(&postBody); err != nil {
		utils.BadRequest(w, "Fail insert data", err.Error())
		return
	}

	validate := validator.New()
	if err := validate.Struct(postBody); err != nil {
		utils.BadRequestErrorFieldEmpty(w, err)
		return
	}

	checkName := config.DB.Where("nationality_name = ?", postBody.NationalityName).First(&nationalities)

	if checkName.RowsAffected > 0 {
		utils.BadRequest(w, "Nationality name already exists", "Error")
		return
	}

	checkCode := config.DB.Where("nationality_code = ?", postBody.NationalityCode).First(&nationalities)

	if checkCode.RowsAffected > 0 {
		utils.BadRequest(w, "Nationality code already exists", "Error")
		return
	}

	data, err := database.InsertNationality(postBody)
	if err != nil {
		utils.InternalServerError(w, err.Error())
	}

	utils.StatusOK(w, 201, "Successfully create nationality!", data)
}
