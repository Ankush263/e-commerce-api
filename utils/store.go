package utils

import (
	"fmt"

	"github.com/ankush263/e-commerce-api/common"
	Models "github.com/ankush263/e-commerce-api/models"
)

type CreateStoreResponse struct {
    Status int `json:"status"`
    Error error `json:"error"`
    Data *Models.StoreResponseModel `json:"data"`
}

type StoreResponse struct {
	Status int `json:"status"`
	Error error `json:"error"`
	Data *[]Models.StoreResponseModel `json:"data"`
}


func CreateStoreInDB(store Models.StoreModel, userid int, storeid string) CreateStoreResponse {
	db := common.SetupDB()
	var response Models.StoreResponseModel

	err := db.QueryRow(`INSERT INTO stores(owner, name, description, store_type, store_id)
						VALUES($1, $2, $3, $4, $5)
						RETURNING id, created_at, updated_at, owner, name, description, store_type, store_id`,
		userid, store.Name, store.Description, store.StoreType, storeid).Scan(
			&response.ID,
			&response.CreatedAt,
			&response.UpdatedAt,
			&response.Owner,
			&response.Name,
			&response.Description,
			&response.StoreType,
			&response.Storeid,
		)

	if err != nil {
		fmt.Println("error: ", err)
		return CreateStoreResponse{
			Status: 0,
			Error: err,
		}
	}

	return CreateStoreResponse{
        Status: 1,
        Data: &response,
    } 
} 

func GetAllStoresFromDB() StoreResponse {
	db := common.SetupDB()

	rows, err := db.Query(`SELECT id, created_at, updated_at, name, owner, description, store_type, store_id FROM stores`)

	if err != nil {
		return StoreResponse{
            Status: 0,
            Error: err,
        }
	}

	defer rows.Close()

	var stores []Models.StoreResponseModel

	for rows.Next() {
		var store Models.StoreResponseModel
		err := rows.Scan(
			&store.ID,
			&store.CreatedAt,	
			&store.UpdatedAt,	
			&store.Name,	
			&store.Owner,	
			&store.Description,	
			&store.StoreType,
			&store.Storeid,
		)
		common.CheckError("Error in getting all stores: ", err)
		stores = append(stores, store)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error: ", err)
		return StoreResponse{
            Status: 0,
            Error: err,
        } 
	}

	return StoreResponse{
		Status: 1,
		Data: &stores,
	}
}

func GetSingleStoreFromDB(id string) CreateStoreResponse {
	db := common.SetupDB()

	var response Models.StoreResponseModel

	err := db.QueryRow(`SELECT id, created_at, updated_at, name, owner, description, store_type, store_id FROM stores WHERE id = $1`, id).Scan(
		&response.ID,
		&response.CreatedAt,
		&response.UpdatedAt,
		&response.Name,
		&response.Owner,
		&response.Description,
		&response.StoreType,
		&response.Storeid,
	)

	if err != nil {
		return CreateStoreResponse{
			Status: 0,
			Error: err,
		}
	}

	return CreateStoreResponse{
		Status: 1,
		Data: &response,
	}
}

func GetStoreByUserid(userid string) CreateStoreResponse {
	db := common.SetupDB()

	var response Models.StoreResponseModel

	err := db.QueryRow(`SELECT id, created_at, updated_at, name, owner, description, store_type, store_id FROM stores WHERE owner = $1`, userid).Scan(
		&response.ID,
		&response.CreatedAt,
		&response.UpdatedAt,
		&response.Name,
		&response.Owner,
		&response.Description,
		&response.StoreType,
		&response.Storeid,
	)

	if err != nil {
		return CreateStoreResponse{
			Status: 0,
			Error: err,
		}
	}

	return CreateStoreResponse{
		Status: 1,
		Data: &response,
	}
}

func GetSingleStoreByStoreId(storeid string) CreateStoreResponse {
	db := common.SetupDB()

	var response Models.StoreResponseModel

	err := db.QueryRow(`SELECT id, created_at, updated_at, name, owner, description, store_type, store_id FROM stores WHERE store_id = $1`, storeid).Scan(
		&response.ID,
		&response.CreatedAt,
		&response.UpdatedAt,
		&response.Name,
		&response.Owner,
		&response.Description,
		&response.StoreType,
		&response.Storeid,
	)

	if err != nil {
		return CreateStoreResponse{
			Status: 0,
			Error: err,
		}
	}

	return CreateStoreResponse{
		Status: 1,
		Data: &response,
	}
}

func UpdateStore(store Models.StoreModel, id string) CreateStoreResponse {
	db := common.SetupDB()

	var response Models.StoreResponseModel
	err := db.QueryRow(`
		UPDATE stores
		SET
			name = COALESCE($1, name),
			description = COALESCE($2, description),
			store_type = COALESCE($3, store_type),
			updated_at = NOW()
		WHERE id = $4
		RETURNING  id, created_at, updated_at, name, owner, description, store_type, store_id;
	`, store.Name, store.Description, store.StoreType, id).Scan(
		&response.ID,
		&response.CreatedAt,
		&response.UpdatedAt,
		&response.Name,
		&response.Owner,
		&response.Description,
		&response.StoreType,
		&response.Storeid,
	)

	if err != nil {
        return CreateStoreResponse{
            Status: 0,
            Error: err,
        }
    }

    return CreateStoreResponse{
        Status: 1,
        Data: &response,
    }
}

func DeleteStore(id string) string {
	db := common.SetupDB()

	result, err := db.Exec(`DELETE FROM stores WHERE id = $1`, id)
	if err != nil {
		fmt.Println("Error: ", err)
		return "error"
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		fmt.Println("Error: ", err)
		return "error"
	}

	if rowsAffected == 0 {
		return "not found"
	}

	return "success"
}
