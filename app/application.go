package app

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rotelando/bookstore_items-api/clients/elasticsearch"
	"net/http"
	"os"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {

	elasticsearch.Init()

	mapUrls()

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	server := &http.Server{
		Handler: loggedRouter,
		Addr:    "localhost:8090",
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
