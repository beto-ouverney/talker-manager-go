package handler

import (
	"net/http"

	userintegration "github.com/beto-ouverney/talker-manager-go/user/integration"
)

//GetUserTokenHandler is a handler for the GetUserToken usecase
func GetUserTokenHandler(w http.ResponseWriter, r *http.Request) {


	token := userintegration.GetUserToken()
	return
}
