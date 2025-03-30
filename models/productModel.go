package models

import "time"

type ProductModel struct {
	Name        *string `json:"name" binding:"required"`
	Description *string `json:"description" binding:"required"`
	Owner       *string `json:"owner"`
	StoreId     *string `json:"store_id"`
	Price       *int    `json:"price" binding:"required"`
}

type ProductResponseModel struct {
	Id          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Owner       string    `json:"owner"`
	StoreId     string    `json:"store_id"`
	Price       int       `json:"price" binding:"required"`
}