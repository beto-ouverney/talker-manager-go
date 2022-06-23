package handler

import (
	"encoding/json"
	"net/http"

	talker "github.com/beto-ouverney/talker-manager-go/talker/entity"
	talkerintegration "github.com/beto-ouverney/talker-manager-go/talker/integration"
)

//AddTalkerHandler is a handler for the POST /talker endpoint.
func AddTalkerHandler(w http.ResponseWriter, r *http.Request) {
	status := 500
	response := []byte("{\"message\":\"Error\"}")
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		var newTalker talker.Talker
		err := decoder.Decode(&newTalker)
		if err != nil {
			errorReturn(w, r, 500, err.Error())
		}

		talkerIntegration := talkerintegration.TalkersIntegration()
		talker, err := talkerIntegration.AddTalker(&newTalker)
		if err != nil {
			errorReturn(w, r, 500, err.Error())
		}

		status = 201
		response = talker

	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}
