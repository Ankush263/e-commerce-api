package utils

import (
	"fmt"

	"github.com/ankush263/e-commerce-api/common"
	Models "github.com/ankush263/e-commerce-api/models"
)

type CreateProductResponse struct {
	Status int `json:"status"`
	Error error `json:"error"`
	Data *Models.ProductResponseModel `json:"data"`
}

type ProductResponse struct {
	Status int `json:"status"`
	Error error `json:"error"`
	Data *[]Models.ProductResponseModel `json:"data"`
}

func CreateProductInDB(product Models.ProductModel, ownerid int, storeid string) CreateProductResponse {
	db := common.SetupDB()
	var response Models.ProductResponseModel

	err := db.QueryRow(`INSERT INTO products(name, description, owner, store_id, price)
						VALUES($1, $2, $3, $4, $5)
						RETURNING id, created_at, updated_at, owner, name, description, price, store_id`,
		product.Name, product.Description, ownerid, storeid, product.Price).Scan(
			&response.Id,
			&response.CreatedAt,
			&response.UpdatedAt,
			&response.Owner,
			&response.Name,
			&response.Description,
			&response.Price,
			&response.StoreId,
		)

	if err != nil {
		fmt.Println("Error: ", err)
		return CreateProductResponse{
			Status: 0,
			Error: err,
		}
	}

	return CreateProductResponse{
		Status: 1,
		Data: &response,
	}
}

func GetAllProducts() ProductResponse {
	db := common.SetupDB()

	rows, err := db.Query(`SELECT id, created_at, updated_at, owner, name, description, price, store_id FROM products`)

	if err != nil {
		return ProductResponse{
			Status: 0,
			Error: err,
		}
	}

	defer rows.Close()

	var products []Models.ProductResponseModel

	for rows.Next() {
		var product Models.ProductResponseModel
		err := rows.Scan(
			&product.Id,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.Owner,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.StoreId,
		)
		common.CheckError("Error in getting all products: ", err)
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error: ", err)
		return ProductResponse{
			Status: 0,
			Error: err,
		}
	}

	return ProductResponse{
		Status: 1,
		Data: &products,
	}

}

func GetSingleProduct(id string) CreateProductResponse {
	db := common.SetupDB()

	var response Models.ProductResponseModel

	err := db.QueryRow(`SELECT id, created_at, updated_at, owner, name, description, price, store_id FROM products WHERE id = $1`, id).Scan(
		&response.Id,
		&response.CreatedAt,
		&response.UpdatedAt,
		&response.Owner,
		&response.Name,
		&response.Description,
		&response.Price,
		&response.StoreId,
	)

	if err != nil {
		return CreateProductResponse{
			Status: 0,
			Error: err,
		}
	}

	return CreateProductResponse{
		Status: 1,
		Data: &response,
	}
}

func GetProductsByStore(storeid string) ProductResponse {
	db := common.SetupDB()

	var products []Models.ProductResponseModel

	rows, err := db.Query(`SELECT id, created_at, updated_at, owner, name, description, price, store_id FROM products WHERE store_id = $1`, storeid)

	if err != nil {
		return ProductResponse{
			Status: 0,
			Error: err,
		}
	}

	for rows.Next() {
		var product Models.ProductResponseModel
		err := rows.Scan(
			&product.Id,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.Owner,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.StoreId,
		)
		common.CheckError("Error in getting all products: ", err)
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error: ", err)
		return ProductResponse{
			Status: 0,
			Error: err,
		}
	}

	return ProductResponse{
		Status: 1,
		Data: &products,
	}
}

func UpdateProduct(product Models.ProductResponseModel, id string) CreateProductResponse {
	db := common.SetupDB()

	var response Models.ProductResponseModel
	err := db.QueryRow(`
		UPDATE products
		SET	
			name = COALESCE($1, name),
			description = COALESCE($2, description),
			price = COALESCE($3, price),
			updated_at = NOW()
		WHERE id = $4
		RETURNING  id, created_at, updated_at, owner, name, description, price, store_id;
	`, product.Name, product.Description, product.Price, id).Scan(
		&response.Id,
		&response.CreatedAt,
		&response.UpdatedAt,
		&response.Owner,
		&response.Name,
		&response.Description,
		&response.Price,
		&response.StoreId,
	)

	if err != nil {
		fmt.Println("Error: ", err)
		return CreateProductResponse{
			Status: 0,
			Error: err,
		}
	}

	return CreateProductResponse{
		Status: 1,
		Data: &response,
	}
}

func DeleteProduct(id string) string {
	db := common.SetupDB()

	result, err := db.Exec(`DELETE FROM products WHERE id = $1`, id)
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
