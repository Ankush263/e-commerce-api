package routers

import (
	"net/http"

	userController "github.com/ankush263/e-commerce-api/controllers"
	"github.com/gorilla/mux"
)

func UserRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/user/signup", userController.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/login", userController.LoginUser).Methods("POST")
	router.Handle("/api/v1/user/profile", userController.AuthMiddleware(http.HandlerFunc(userController.GetProfile)))
	router.HandleFunc("/api/v1/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/user/{id}", userController.GetSingleUser).Methods("GET")
	router.HandleFunc("/api/v1/user/{id}", userController.UpdateSingleUser).Methods("PATCH")
	router.HandleFunc("/api/v1/user/{id}", userController.DeleteUserById).Methods("DELETE")

	return router
}
