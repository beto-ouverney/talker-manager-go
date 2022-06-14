package main

import (
	"net/http"

	"github.com/beto-ouverney/talker-manager-go/handler"
	"github.com/beto-ouverney/talker-manager-go/myrouter"
)

func main() {
	router := &myrouter.Router{}
	router.Route(http.MethodGet, `/talkers/(?P<id>\d+)`, handler.GetTalkerByIDHandler)
	router.Route(http.MethodGet, "/talkers", handler.GetAllTalkersHandler)
	router.Route(http.MethodPost, "/login", handler.GetUserTokenHandler)
	http.ListenAndServe(":8080", router)
}
