package models

import "time"

type UsersModel struct {
	UserName *string `json:"username" binding:"required,alphanum"`
	Email    *string `json:"email" binding:"required,email"`
	Password *string `json:"password" binding:"required,min=8,max=20"`
	Phone    *string `json:"phone" binding:"required,min=10,max=12"`
}

type ResponseUsersModel struct {
	ID        int    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserName  string `json:"username" binding:"required,alphanum"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=20"`
	Phone     string `json:"phone" binding:"required,min=10,max=12"`
}

type ResponseDBModel struct {
	Status int `json:"status"`
	Data *ResponseUsersModel `json:"data,omitempty"`
}

type ResponseModel struct {
	Status int   `json:"status"`
	Message string `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
