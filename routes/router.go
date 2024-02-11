package routes

import (
	"github.com/gorilla/mux"
	"github.com/sandeepkumarnayak/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/addproduct", controllers.AddProduct).Methods("POST")
	router.HandleFunc("/api/products", controllers.GetAllProducts).Methods("GET")
	router.HandleFunc("/api/product/{id}", controllers.GetProductByID).Methods("GET")
	router.HandleFunc("/api/product/{id}", controllers.DeleteProductById).Methods("DELETE")
	return router
}
