package app

import "github.com/gorilla/mux"

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()
}
