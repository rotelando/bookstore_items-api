package app

import (
	"github.com/gorilla/handlers"
	"github.com/rotelando/bookstore_items-api/controllers"
	"net/http"
	"os"
)

func mapUrls() {

	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items", controllers.ItemController.Create).Methods(http.MethodPost)

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	server := &http.Server{
		Handler: loggedRouter,
		Addr:    "localhost:8090",
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
