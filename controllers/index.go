package controllers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status int `json:"status"`
	Msg string `json:"message"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Response{
		Status: 1,
		Msg: "Hello from GO",
	})
}
