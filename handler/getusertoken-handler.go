package handler

import (
	"encoding/json"
	"net/http"

	userentity "github.com/beto-ouverney/talker-manager-go/user/entity"
	userintegration "github.com/beto-ouverney/talker-manager-go/user/integration"
)

//GetUserTokenHandler is a handler for the GetUserToken usecase
func GetUserTokenHandler(w http.ResponseWriter, r *http.Request) {
	status := 500
	response := []byte("{\"message\":\"Error\"}")
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		var user userentity.User
		err := decoder.Decode(&user)
		if err != nil {
			errorReturn(w, r, 500, err.Error())
		}
		userIntegration := userintegration.UserIntegration()
		token := userIntegration.GetUserToken(user)
		status = 200
		tokenStruct := struct {
			Token string `json:"token"`
		}{
			Token: *token,
		}
		response, err = json.Marshal(&tokenStruct)
		if err != nil {
			errorReturn(w, r, 500, err.Error())
		}
	}
	_, err := w.Write(response)
	if err != nil {
		errorReturn(w, r, 500, err.Error())
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return
}
