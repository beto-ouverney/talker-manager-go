package handler

import (
	"net/http"

	talkerintegration "github.com/beto-ouverney/talker-manager-go/talker/integration"
)

//GetAllTalkersHandler is a handler for the GET /talkers endpoint.
func GetAllTalkersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		talkerIntegration := talkerintegration.TalkersIntegration()
		talkers, err := talkerIntegration.GetAllTalkers()
		if err != nil {
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(talkers))

	}
}
