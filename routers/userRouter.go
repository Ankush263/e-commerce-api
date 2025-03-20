package routers

import (
	userController "github.com/ankush263/e-commerce-api/controllers"
	"github.com/gorilla/mux"
)

func UserRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/users", userController.GetUsers).Methods("GET")

	return router
}
