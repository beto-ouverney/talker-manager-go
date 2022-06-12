package handler

import (
	"net/http"
	"strconv"
	"encoding/json" 
 
	talkerintegration "github.com/beto-ouverney/talker-manager-go/talker/integration"
)

//GetTalkerByIDHandler is a handler for the GET /talkers/:id endpoint.
func GetTalkerByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		talkerIntegration := talkerintegration.TalkersIntegration()
		id,err := strconv.Atoi(URLParam(r, "id"))
		talker, err := talkerIntegration.GetTalkerByID(id)
		if err != nil {
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		talkerJSON, err := json.Marshal(talker)
		w.Write(talkerJSON)
	}
}
