package main

import (
	"go-mux-gorm-example/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	routes.UserRoutes(router)

	log.Fatal(http.ListenAndServe("localhost:3000", router))
}
