package routes

import (
	"github.com/gorilla/mux"
	"github.com/rrscodes/fashionfever-api/pkg/controllers"
)

var RegisterDressRoutes = func(router *mux.Router) {
	router.HandleFunc("/dress/", controllers.CreateDress).Methods("POST")
	router.HandleFunc("/dress/", controllers.GetAllDresses).Methods("GET")
	router.HandleFunc("/dress/{dressId}", controllers.GetDressById).Methods("GET")
	router.HandleFunc("/dress/{dressId}", controllers.UpdateDress).Methods("PUT")
	router.HandleFunc("/dress/{dressId}", controllers.DeleteDress).Methods("DELETE")
}
