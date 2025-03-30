package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ankush263/e-commerce-api/models"
	"github.com/ankush263/e-commerce-api/utils"
	"github.com/gorilla/mux"
)


func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var reqProduct models.ProductModel
	_ = json.NewDecoder(r.Body).Decode(&reqProduct)

	// Check request body
	if reqProduct.Name == nil {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 0,
			Message: "Product name is missing",
		})
		return
	}
	if reqProduct.Description == nil {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 0,
			Message: "Product description is missing",
		})
		return
	}
	if reqProduct.Price == nil {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 0,
			Message: "Product price is missing",
		})
		return
	}

	userId := r.Context().Value("user_id")
	
	// storeid := uuid.New().String()
	var storeid string
	
	strUserid, ok := userId.(string)
	
	//* Check if User type is seller or not (Only seller can create stores)
	userResp := utils.GetSingleUserInDB(strUserid)

	if userResp.Status == 0 {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 0,
			Message: "User with this id does not exists",
		})
	} else {
		userData := userResp.Data

		if userData.Role != "seller" {
			json.NewEncoder(w).Encode(models.ResponseModel{
				Status: 0,
				Message: "Only Seller can create Products",
			})
			return
		} else {
			storeResp := utils.GetStoreByUserid(strUserid)

			if storeResp.Status == 0 {
				json.NewEncoder(w).Encode(models.ResponseModel{
					Status: 0,
					Message: "Can't find a store",
				})
				return
			}

			storeid = storeResp.Data.Storeid
		}
	}


	if !ok {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status:  0,
			Message: "Invalid user ID",
			Data:    nil,
		})
		return
	}

	intUserid, err := strconv.Atoi(strUserid)

	if err != nil {
		fmt.Println("error: userId is not an integer")
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status:  0,
			Message: "Invalid user ID",
			Data:    nil,
		})
		return
	}

	response := utils.CreateProductInDB(reqProduct, intUserid, storeid)

	if response.Status == 0 {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 0,
			Message: "Failed",
			Error: response.Error,
		})
	} else {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 1,
			Message: "Success",
			Data: response.Data,
		})
	}
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	allStores := utils.GetAllProducts()

	if allStores.Status == 0 {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 0,
			Message: "Products not found",
			Data: nil,
		})
	} else {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 1,
			Message: "Success",
			Data: allStores.Data,
		})
	}
}

func GetSingleProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	storedata := utils.GetSingleProduct(id)

	if storedata.Status == 0 {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 0,
			Message: "Product with that id does not exist",
			Data: nil,
		})
	} else {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 1,
			Message: "Success",
			Data: storedata.Data,
		})
	}
}

func GetProductByStoreid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["storeid"]

	storedata := utils.GetProductsByStore(id)

	if storedata.Status == 0 {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 0,
			Message: "Products with that storeid does not exist",
			Data: nil,
		})
	} else {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 1,
			Message: "Success",
			Data: storedata.Data,
		})
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PATCH")

	var reqProduct models.ProductResponseModel
	_ = json.NewDecoder(r.Body).Decode(&reqProduct)

	params := mux.Vars(r)
	id := params["id"]

	userId := r.Context().Value("user_id")
	// storeid := uuid.New().String()

	strUserid, ok := userId.(string)

	if !ok {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status:  0,
			Message: "Invalid user ID",
			Data:    nil,
		})
		return
	}

	intUserid, err := strconv.Atoi(strUserid)

	if err != nil {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status:  0,
			Message: "Invalid user ID",
			Data:    nil,
		})
		return
	}
	
	//* Check if User type is seller or not 
	userResp := utils.GetSingleUserInDB(strUserid)

	if userResp.Status == 0 {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 0,
			Message: "User with this id does not exists",
		})
		return
	} else {
		userData := userResp.Data

		if userData.Role != "seller" {
			json.NewEncoder(w).Encode(models.ResponseModel{
				Status: 0,
				Message: "Only Seller can update Products",
			})
			return
		} else {
			if userData.ID != intUserid {
				json.NewEncoder(w).Encode(models.ResponseModel{
					Status: 0,
					Message: "Only owner of the product can update it.",
				})
				return
			}
		}
	}

	updatedData := utils.UpdateProduct(reqProduct, id)

	if updatedData.Status == 0 {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 0,
			Message: "Product updates failed",
		})
	} else {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 1,
			Data: updatedData.Data,
		})
	}

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	userId := r.Context().Value("user_id")
	
	strUserid, ok := userId.(string)

	if !ok {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status:  0,
			Message: "Invalid user ID",
			Data:    nil,
		})
		return
	}

	intUserid, err := strconv.Atoi(strUserid)

	if err != nil {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status:  0,
			Message: "Invalid user ID",
			Data:    nil,
		})
		return
	}
	
	//* Check if User type is seller or not 
	userResp := utils.GetSingleUserInDB(strUserid)

	if userResp.Status == 0 {
		json.NewEncoder(w).Encode(models.ResponseModel{
			Status: 0,
			Message: "User with this id does not exists",
		})
		return
	} else {
		userData := userResp.Data

		if userData.Role != "seller" {
			json.NewEncoder(w).Encode(models.ResponseModel{
				Status: 0,
				Message: "Only Seller can delete Products",
			})
			return
		} else {
			if userData.ID != intUserid {
				json.NewEncoder(w).Encode(models.ResponseModel{
					Status: 0,
					Message: "Only owner of the product can delete it.",
				})
				return
			}
		}
	}

	deleteProduct := utils.DeleteProduct(id)

	if deleteProduct == "success" {
		json.NewEncoder(w).Encode("Deleted")
	} else {
		json.NewEncoder(w).Encode("Failed to delete the store")
	}
}
