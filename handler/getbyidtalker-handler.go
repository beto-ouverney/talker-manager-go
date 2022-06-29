package handler

import (
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
			if err != nil {
				errorReturn(w, r, 500, err.Error())
			}
			response = talker
		} else {
			status = 404

			response = []byte("{\"message\":\"Talker not found\"}")
		}
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
