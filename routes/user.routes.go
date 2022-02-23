package routes

import (
	"go-mux-gorm-example/controllers"

	"github.com/gorilla/mux"
)

func UserRoutes(
	router *mux.Router,
) {
	router.HandleFunc("/users", controllers.GetAllUSers()).Methods("GET")
}
