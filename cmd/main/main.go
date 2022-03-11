package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rrscodes/fashionfever-api/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterDressRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("Localhost:9010", r))
}
