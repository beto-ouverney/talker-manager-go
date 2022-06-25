package main

import (
	"net/http"

	"github.com/beto-ouverney/talker-manager-go/handler"
	"github.com/beto-ouverney/talker-manager-go/middleware"
	"github.com/beto-ouverney/talker-manager-go/myrouter"
)

func main() {
	router := &myrouter.Router{}
	router.Route(http.MethodGet, `/talker/search`, []myrouter.Middleware{middleware.TokenValidate}, handler.SearchTalkersHandler)
	router.Route(http.MethodGet, `/talkers/(?P<id>\d+)`, nil, handler.GetTalkerByIDHandler)
	router.Route(http.MethodGet, "/talkers", nil, handler.GetAllTalkersHandler)
	router.Route(http.MethodPost, "/login", []myrouter.Middleware{middleware.UserValidate}, handler.GetUserTokenHandler)
	router.Route(http.MethodPut, `/talker/(?P<id>\d+)`, []myrouter.Middleware{middleware.TokenValidate, middleware.TalkerValidate}, handler.EditTalkerHandler)
	router.Route(http.MethodDelete, `/talker/(?P<id>\d+)`, []myrouter.Middleware{middleware.TokenValidate}, handler.DeleteTalkerHandler)
	router.Route(http.MethodPost, "/talker", []myrouter.Middleware{middleware.TokenValidate, middleware.TalkerValidate}, handler.AddTalkerHandler)
	http.ListenAndServe(":8080", router)
}
