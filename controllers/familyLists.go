package controllers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"strconv"
	"test-bookingtogo/lib/database"
	"test-bookingtogo/lib/utils"
	"test-bookingtogo/models"
)

func GetFamilyLists(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	offset, limit, page := utils.GetPagination(pageStr, limitStr)

	data, totalData, err := database.GetFamilyLists(offset, limit)
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))

	response := models.ResponseGetList{
		Code:       200,
		Message:    "Successfully get family lists!",
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

func GetFamilyListById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strId := vars["id"]

	id, err := strconv.Atoi(strId)
	if err != nil {
		utils.BadRequest(w, "Failed to get id!", "Error")
		return
	}

	data, err := database.GetFamilyListById(id)
	if err != nil {
		utils.NotFound(w, err.Error(), "Not Found")
		return
	}
	response := models.Response{
		Code:    200,
		Message: "Successfully get family list by id!",
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if errJson := json.NewEncoder(w).Encode(response); errJson != nil {
		utils.BadRequest(w, "Failed to encode response", "Error")
		return
	}
}

func CreateFamilyList(w http.ResponseWriter, r *http.Request) {
	var postBody models.FamilyListPost

	if err := json.NewDecoder(r.Body).Decode(&postBody); err != nil {
		utils.BadRequest(w, "Fail insert data", err.Error())
		return
	}

	validate := validator.New()
	if errValidate := validate.Struct(postBody); errValidate != nil {
		utils.BadRequestErrorFieldEmpty(w, errValidate)
		return
	}

	_, errGetCustomer := database.GetCustomerById(postBody.CstId)
	if errGetCustomer != nil {
		utils.NotFound(w, errGetCustomer.Error(), "Not Found")
		return
	}

	data, err := database.InsertFamilyList(postBody)
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}

	response := models.Response{
		Code:    201,
		Message: "Successfully add family list!",
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if errJson := json.NewEncoder(w).Encode(response); errJson != nil {
		utils.BadRequest(w, "Failed to encode response", "Error")
		return
	}
}

func UpdateFamilyList(w http.ResponseWriter, r *http.Request) {
	var putBody models.FamilyListUpdate
	vars := mux.Vars(r)
	strId := vars["id"]

	id, err := strconv.Atoi(strId)
	if err != nil {
		utils.BadRequest(w, "Failed to get id!", "Error")
		return
	}

	_, err = database.GetFamilyListById(id)
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

	_, errGetCustomer := database.GetCustomerById(putBody.CstId)
	if errGetCustomer != nil {
		utils.NotFound(w, errGetCustomer.Error(), "Not Found")
		return
	}

	data, err := database.UpdateFamilyList(id, putBody)
	if err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}

	response := models.Response{
		Code:    201,
		Message: "Successfully update family list!",
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if errJson := json.NewEncoder(w).Encode(response); errJson != nil {
		utils.BadRequest(w, "Failed to encode response", "Error")
		return
	}
}

func DeleteFamilyList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strId := vars["id"]

	id, err := strconv.Atoi(strId)

	_, err = database.GetFamilyListById(id)
	if err != nil {
		utils.NotFound(w, err.Error(), "Not Found")
		return
	}

	if err := database.DeleteFamilyList(id); err != nil {
		utils.InternalServerError(w, err.Error())
		return
	}

	response := map[string]interface{}{
		"Code":    http.StatusOK,
		"Message": "Successfully delete family list!",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if errJson := json.NewEncoder(w).Encode(response); errJson != nil {
		utils.BadRequest(w, "Failed to encode response", "Error")
		return
	}
}
