package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	talker "github.com/beto-ouverney/talker-manager-go/talker/entity"
	talkerintegration "github.com/beto-ouverney/talker-manager-go/talker/integration"
)

//EditTalkerHandler is a handler for the PUT /talker/{id} endpoint.
func EditTalkerHandler(w http.ResponseWriter, r *http.Request) {
	status := 500
	response := []byte("{\"message\":\"Error\"}")
	if r.Method == http.MethodPut {
		decoder := json.NewDecoder(r.Body)
		var newTalker talker.Talker
		err := decoder.Decode(&newTalker)
		if err == nil {
			id, err := strconv.Atoi(URLParam(r, "id"))
			if err == nil {
				newTalker.ID = id
				if err == nil {
					talkerIntegration := talkerintegration.TalkersIntegration()
					talker, err := talkerIntegration.EditTalker(&newTalker)
					if err != nil {
						errorReturn(w, r, 500, err.Error())
					}
					status = 200
					response = talker
				}
			}
		}
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
