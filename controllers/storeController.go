package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	model "github.com/ankush263/e-commerce-api/models"
	"github.com/ankush263/e-commerce-api/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var reqStore model.StoreModel
	_ = json.NewDecoder(r.Body).Decode(&reqStore)

	userId := r.Context().Value("user_id")

	storeid := uuid.New().String()

	strUserid, ok := userId.(string)

	if !ok {
		fmt.Println("error: userId is not an string")
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status:  0,
			Message: "Invalid user ID",
			Data:    nil,
		})
		return
	}

	intUserid, err := strconv.Atoi(strUserid)

	if err != nil {
		fmt.Println("error: userId is not an integer")
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status:  0,
			Message: "Invalid user ID",
			Data:    nil,
		})
		return
	}

	response := utils.CreateStoreInDB(reqStore, intUserid, storeid)

	if response.Status == 0 {
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status: 0,
			Message: "Failed",
			Error: response.Error,
		})
	} else {
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status: 1,
			Message: "Success",
			Data: response.Data,
		})
	}
}

func GetAllStores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	allStores := utils.GetAllStoresFromDB()

	if allStores.Status == 0 {
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status: 0,
			Message: "Stores not found",
			Data: nil,
		})
	} else {
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status: 0,
			Message: "Success",
			Data: allStores.Data,
		})
	}
}

func GetSingleStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	storedata := utils.GetSingleStoreFromDB(id)

	if storedata.Status == 0 {
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status: 0,
			Message: "Store with that id does not exist",
			Data: nil,
		})
	} else {
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status: 0,
			Message: "Success",
			Data: storedata.Data,
		})
	}
}
