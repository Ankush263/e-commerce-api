package routers

import (
	"log"
	"net/http"

	controller "github.com/ankush263/e-commerce-api/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	log.Println("Initializing Router...")
	router := mux.NewRouter().StrictSlash(true)
	
	router.HandleFunc("/api/v1", controller.Index).Methods("GET")

	router.HandleFunc("/api/v1/user/signup", controller.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/login", controller.LoginUser).Methods("POST")
	router.Handle("/api/v1/user/profile", controller.AuthMiddleware(http.HandlerFunc(controller.GetProfile)))
	router.HandleFunc("/api/v1/users", controller.GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/user/{id}", controller.GetSingleUser).Methods("GET")
	router.HandleFunc("/api/v1/user/{id}", controller.UpdateSingleUser).Methods("PATCH")
	router.HandleFunc("/api/v1/user/{id}", controller.DeleteUserById).Methods("DELETE")

	router.Handle("/api/v1/store/create", controller.AuthMiddleware(http.HandlerFunc(controller.CreateStore))).Methods("POST")
	router.HandleFunc("/api/v1/store", controller.GetAllStores).Methods("GET")
	router.HandleFunc("/api/v1/store/{id}", controller.GetSingleStore).Methods("GET")
	router.Handle("/api/v1/store/{id}", controller.AuthMiddleware(http.HandlerFunc(controller.UpdateStore))).Methods("PATCH")
	router.Handle("/api/v1/store/{id}", controller.AuthMiddleware(http.HandlerFunc(controller.DeleteStore))).Methods("DELETE")
	router.HandleFunc("/api/v1/store/id/{storeid}", controller.GetStoreByStoreid).Methods("GET")

	router.Handle("/api/v1/product/create", controller.AuthMiddleware(http.HandlerFunc(controller.CreateProduct))).Methods("POST")
	router.HandleFunc("/api/v1/product", controller.GetAllProducts).Methods("GET")
	router.HandleFunc("/api/v1/product/{id}", controller.GetSingleProduct).Methods("GET")
	router.Handle("/api/v1/product/{id}", controller.AuthMiddleware(http.HandlerFunc(controller.UpdateProduct))).Methods("PATCH")
	router.Handle("/api/v1/product/{id}", controller.AuthMiddleware(http.HandlerFunc(controller.DeleteProduct))).Methods("DELETE")
	router.HandleFunc("/api/v1/product/id/{storeid}", controller.GetProductByStoreid).Methods("GET")

	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err == nil {
			log.Println("Registered Route:", path)
		}
		return nil
	})
	


	return router
}
