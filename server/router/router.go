package router

import (
	"productManagement/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// router.HandleFunc("/users", controller.GetAllTheUsers).Methods("GET")
	router.HandleFunc("/allproducts", controller.GetAllTheProducts).Methods("GET")
	router.HandleFunc("/register", controller.SignUpUser).Methods("POST")
	router.HandleFunc("/login", controller.SignInUser).Methods("POST") //We have to check this
	router.HandleFunc("/add/product", controller.AddOneProduct).Methods("POST")
	router.HandleFunc("/update/product/{id}", controller.EditOneProduct).Methods("PUT")
	router.HandleFunc("/delete/product/{id}", controller.DeleteAProduct).Methods("DELETE")
	return router

}
