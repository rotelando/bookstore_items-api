package app

import (
	"github.com/rotelando/bookstore_items-api/controllers"
	"net/http"
)

func mapUrls() {

	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)

	router.HandleFunc("/items", controllers.ItemController.Create).Methods(http.MethodPost)
}
