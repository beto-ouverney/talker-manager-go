package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	talkerintegration "github.com/beto-ouverney/talker-manager-go/talker/integration"
)

//GetTalkerByIDHandler is a handler for the GET /talkers/:id endpoint.
func GetTalkerByIDHandler(w http.ResponseWriter, r *http.Request) {
	status := 500
	response := []byte("{\"message\":\"Error\"}")
	if r.Method == http.MethodGet {
		talkerIntegration := talkerintegration.TalkersIntegration()
		id, err := strconv.Atoi(URLParam(r, "id"))
		talker, err := talkerIntegration.GetTalkerByID(id)
		if err != nil {
			errorReturn(w, r, 500, err.Error())
		}
		if talker != nil {
			status = 200
			talkerJSON, err := json.Marshal(talker)
			if err != nil {
				errorReturn(w, r, 500, err.Error())
			}
			response = talkerJSON
		} else {
			status = 404

			w.Write([]byte("{\"message\":\"Talker not found\"}"))
		}
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
