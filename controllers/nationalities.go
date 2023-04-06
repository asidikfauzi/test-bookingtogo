package controllers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"strconv"
	"test-bookingtogo/config"
	"test-bookingtogo/lib/database"
	"test-bookingtogo/lib/utils"
	"test-bookingtogo/models"
)

var nationalities []models.Nationalities

func GetNationalities(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	offset, limit, page := utils.GetPagination(pageStr, limitStr)

	data, totalData, err := database.GetNationalities(offset, limit)
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))

	response := models.ResponseGetList{
		200,
		"Successfully get nationalities!",
		data,
		page,
		limit,
		totalPage,
		totalData,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.BadRequest(w, "Failed to encode response", "Error")
		return
	}
}

func GetNationalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strId := vars["id"]

	id, err := strconv.Atoi(strId)
	if err != nil {
		utils.BadRequest(w, "Failed to get id!", "Error")
		return
	}

	data, err := database.GetNationalityById(id)
	if err != nil {
		utils.NotFound(w, err.Error(), "Not Found")
		return
	}
	response := models.Response{
		200,
		"Successfully get nationality by id!",
		data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.BadRequest(w, "Failed to encode response", "Error")
		return
	}
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
		return
	}

	response := models.Response{
		201,
		"Successfully add nationality!",
		data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.BadRequest(w, "Failed to encode response", "Error")
		return
	}
}

func UpdateNationality(w http.ResponseWriter, r *http.Request) {
	var putBody models.NationalityUpdate
	vars := mux.Vars(r)
	strId := vars["id"]

	id, err := strconv.Atoi(strId)

	_, err = database.GetNationalityById(id)
	if err != nil {
		utils.NotFound(w, err.Error(), "Not Found")
		return
	}

	if err != nil {
		utils.BadRequest(w, "Failed to get id!", "Error")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&putBody); err != nil {
		utils.BadRequest(w, "Fail insert data", err.Error())
		return
	}

	validate := validator.New()
	if err := validate.Struct(putBody); err != nil {
		utils.BadRequestErrorFieldEmpty(w, err)
		return
	}

	checkName := config.DB.Where("nationality_name = ? AND nationality_id != ?",
		putBody.NationalityName, id).Find(&nationalities)
	if checkName.RowsAffected > 0 {
		utils.BadRequest(w, "Nationality name already exists", "Error")
		return
	}

	checkCode := config.DB.Where("nationality_code = ? AND nationality_id != ?",
		putBody.NationalityCode, id).Find(&nationalities)
	if checkCode.RowsAffected > 0 {
		utils.BadRequest(w, "Nationality code already exists", "Error")
		return
	}

	data, err := database.UpdateNationality(id, putBody)
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}

	response := models.Response{
		201,
		"Successfully update nationality!",
		data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.BadRequest(w, "Failed to encode response", "Error")
		return
	}
}
