package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"test-bookingtogo/config"
	"test-bookingtogo/lib/database"
	"test-bookingtogo/lib/utils"
	"test-bookingtogo/models"
)

var customers []models.Customers

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	offset, limit, page := utils.GetPagination(pageStr, limitStr)

	data, totalData, err := database.GetCustomers(offset, limit)
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))

	response := models.ResponseGetList{
		Code:       200,
		Message:    "Successfully get customers!",
		Data:       data,
		Page:       page,
		Limit:      limit,
		TotalPage:  totalPage,
		TotalCount: totalData,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if errJson := json.NewEncoder(w).Encode(response); errJson != nil {
		utils.BadRequest(w, "Failed to encode response", "Error")
		return
	}
}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strId := vars["id"]

	id, err := strconv.Atoi(strId)
	if err != nil {
		utils.BadRequest(w, "Failed to get id!", "Error")
		return
	}

	data, err := database.GetCustomerById(id)
	if err != nil {
		utils.NotFound(w, err.Error(), "Not Found")
		return
	}
	response := models.Response{
		Code:    200,
		Message: "Successfully get customers by id!",
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if errJson := json.NewEncoder(w).Encode(response); errJson != nil {
		utils.BadRequest(w, "Failed to encode response", "Error")
		return
	}
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var postBody models.CustomerPost

	if err := json.NewDecoder(r.Body).Decode(&postBody); err != nil {
		utils.BadRequest(w, "Fail insert data", err.Error())
		return
	}

	validate := validator.New()
	if errValidate := validate.Struct(postBody); errValidate != nil {
		utils.BadRequestErrorFieldEmpty(w, errValidate)
		return
	}

	if !emailRegex.MatchString(postBody.CstEmail) {
		utils.BadRequest(w, "Invalid email format", "Bad Request")
		return
	}

	_, errGetNationality := database.GetNationalityById(postBody.NationalityId)
	fmt.Println(errGetNationality)
	if errGetNationality != nil {
		utils.NotFound(w, errGetNationality.Error(), "Not Found")
		return
	}

	checkEmail := config.DB.Where("cst_email = ?", postBody.CstEmail).First(&customers)
	if checkEmail.RowsAffected > 0 {
		utils.BadRequest(w, "Email already exists", "Error")
		return
	}

	data, err := database.InsertCustomer(postBody)
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}

	response := models.Response{
		Code:    201,
		Message: "Successfully add customer!",
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if errJson := json.NewEncoder(w).Encode(response); errJson != nil {
		utils.BadRequest(w, "Failed to encode response", "Error")
		return
	}
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var putBody models.CustomerUpdate
	vars := mux.Vars(r)
	strId := vars["id"]

	id, err := strconv.Atoi(strId)
	if err != nil {
		utils.BadRequest(w, "Failed to get id!", "Error")
		return
	}

	_, err = database.GetCustomerById(id)
	if err != nil {
		utils.NotFound(w, err.Error(), "Not Found")
		return
	}

	if errJson := json.NewDecoder(r.Body).Decode(&putBody); errJson != nil {
		utils.BadRequest(w, "Fail insert data", errJson.Error())
		return
	}

	validate := validator.New()
	if errValidate := validate.Struct(putBody); errValidate != nil {
		utils.BadRequestErrorFieldEmpty(w, errValidate)
		return
	}

	if !emailRegex.MatchString(putBody.CstEmail) {
		utils.BadRequest(w, "Invalid email format", "Bad Request")
		return
	}

	_, errGetNationality := database.GetNationalityById(putBody.NationalityId)
	if errGetNationality != nil {
		utils.NotFound(w, errGetNationality.Error(), "Not Found")
		return
	}

	checkEmail := config.DB.Where("cst_email = ? AND cst_id != ?",
		putBody.CstEmail, id).Find(&customers)
	if checkEmail.RowsAffected > 0 {
		utils.BadRequest(w, "Email already exists", "Error")
		return
	}

	data, err := database.UpdateCustomer(id, putBody)
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}

	response := models.Response{
		Code:    201,
		Message: "Successfully update customers!",
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if errJson := json.NewEncoder(w).Encode(response); errJson != nil {
		utils.BadRequest(w, "Failed to encode response", "Error")
		return
	}
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strId := vars["id"]

	id, err := strconv.Atoi(strId)

	_, err = database.GetCustomerById(id)
	if err != nil {
		utils.NotFound(w, err.Error(), "Not Found")
		return
	}

	if err := database.DeleteCustomer(id); err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}

	response := map[string]interface{}{
		"Code":    http.StatusOK,
		"Message": "Successfully delete customers!",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if errJson := json.NewEncoder(w).Encode(response); errJson != nil {
		utils.BadRequest(w, "Failed to encode response", "Error")
		return
	}
}
