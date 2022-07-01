package handler

import (
	"net/http"
	"strconv"

	talkerintegration "github.com/beto-ouverney/talker-manager-go/talker/integration"
)

//DeleteTalkerHandler is a handler for the DELETE /talkers/:id endpoint.
func DeleteTalkerHandler(w http.ResponseWriter, r *http.Request) {
	status := 500
	response := []byte("{\"message\":\"Error\"}")
	if r.Method == http.MethodDelete {
		id, err := strconv.Atoi(URLParam(r, "id"))
		if err == nil {
			talkerIntegration := talkerintegration.TalkersIntegration()
			err := talkerIntegration.DeleteTalker(id)
			if err == nil {
				status = 204
				response = nil

			} else if err.Error() == "Talker not found" {
				status = 404
				response = []byte("{\"message\":\"Talker not found\"}")
			} else {
				errorReturn(w, r, 500, err.Error())
			}
		}

	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
