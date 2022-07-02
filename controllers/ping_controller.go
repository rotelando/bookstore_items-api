package controllers

import "net/http"

var (
	PingController pingControllerInterface = &pingController{}
	pong                                   = "pong"
)

type pingControllerInterface interface {
	Ping(http.ResponseWriter, *http.Request)
}

type pingController struct{}

func (p *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pong))
	return
}
