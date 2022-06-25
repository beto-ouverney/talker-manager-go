package handler

import (
	"net/http"

	talkerintegration "github.com/beto-ouverney/talker-manager-go/talker/integration"
)

//SearchTalkersHandler is a handler for the search talkers
func SearchTalkersHandler(w http.ResponseWriter, r *http.Request) {
	status := 500
	response := []byte("{\"message\":\"Error\"}")
	if r.Method == http.MethodGet {
		query := r.URL.Query().Get("q")
		talkerIntegration := talkerintegration.TalkersIntegration()
		talkers, err := talkerIntegration.SearchTalkers(query)
		if err == nil {
			status = 200
			response = talkers
		}
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
