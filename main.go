package main

import (
	"net/http"

	"github.com/beto-ouverney/talker-manager-go/handler"
)

func main() {
	http.HandleFunc("/talkers", handler.GetAllTalkersHandler)
	http.ListenAndServe(":8080", nil)
}
