package models

import "time"

type StoreModel struct {
	Owner       *int      `json:"owner"`
	Name        *string   `json:"name" binding:"required,alphanum"`
	Description *string   `json:"description"`
	StoreType   *string `json:"store_type" binding:"required"`
	Storeid     *string      `json:"storeid" binding:"required"`
}

type StoreResponseModel struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Owner       int       `json:"owner"`
	Name        string    `json:"name" binding:"required,alphanum"`
	Description string    `json:"description"`
	StoreType   string  `json:"store_type" binding:"required"`
	Storeid     string       `json:"storeid" binding:"required"`
}
