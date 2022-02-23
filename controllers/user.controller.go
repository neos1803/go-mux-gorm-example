package controllers

import (
	"encoding/json"
	"go-mux-gorm-example/models"
	"net/http"
)

func GetAllUSers() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(models.MockUsers())
	}
}
