package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ankush263/e-commerce-api/models"
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

	// Check request body
	if reqStore.Name == nil {
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status: 0,
			Message: "Store name is missing",
		})
		return
	}
	if reqStore.Description == nil {
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status: 0,
			Message: "Store description is missing",
		})
		return
	}
	if reqStore.StoreType == nil {
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status: 0,
			Message: "Store type is missing",
		})
		return
	}

	userId := r.Context().Value("user_id")
	
	storeid := uuid.New().String()
	
	strUserid, ok := userId.(string)
	
	//* Check if User type is seller or not (Only seller can create stores)
	userResp := utils.GetSingleUserInDB(strUserid)

	if userResp.Status == 0 {
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status: 0,
			Message: "User with this id does not exists",
		})
	} else {
		userData := userResp.Data

		if userData.Role != "seller" {
			json.NewEncoder(w).Encode(model.ResponseModel{
				Status: 0,
				Message: "Only Seller can create stores",
			})
			return
		}
	}

	if !ok {
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
			Status: 1,
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
			Status: 1,
			Message: "Success",
			Data: storedata.Data,
		})
	}
}

func GetStoreByStoreid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["storeid"]

	storedata := utils.GetSingleStoreByStoreId(id)

	if storedata.Status == 0 {
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status: 0,
			Message: "Store with that store id does not exist",
			Data: nil,
		})
	} else {
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status: 1,
			Message: "Success",
			Data: storedata.Data,
		})
	}
}

func CheckOwner(storeid string, userId string) bool {
	response := utils.GetSingleStoreFromDB(storeid)

	if response.Status == 0 {
		return false
	} else {
		intId, err := strconv.Atoi(userId)
		if err != nil {
			return false
		}
		if response.Data.ID == intId  {
			return true
		}
	}
	return false
}

func UpdateStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PATCH")

	var reqStore model.StoreModel
	_ = json.NewDecoder(r.Body).Decode(&reqStore)

	params := mux.Vars(r)
	id := params["id"]

	userId := r.Context().Value("user_id").(string)

	//* Check for the owner of the store
	isOwner := CheckOwner(id, userId)

	if !isOwner {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 0,
			Message: "Only owner of the store can update the store details.",
		})
		return
	}

	updatedData := utils.UpdateStore(reqStore, id)

	if updatedData.Status == 0 {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 0,
			Message: "Store updates failed",
		})
	} else {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 1,
			Data: updatedData.Data,
		})
	}

}

func DeleteStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	userId := r.Context().Value("user_id").(string)

	//* Check for the owner of the store
	isOwner := CheckOwner(id, userId)

	if !isOwner {
		json.NewEncoder(w).Encode(model.ResponseModel{
			Status: 0,
			Message: "Only owner of the store can delete the store.",
		})
		return
	}

	deleteStore := utils.DeleteStore(id)

	if deleteStore == "success" {
		json.NewEncoder(w).Encode("Deleted")
	} else {
		json.NewEncoder(w).Encode("Failed to delete the store")
	}
}
