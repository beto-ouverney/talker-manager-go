package middleware

import (
	"encoding/json"
	"regexp"

	errorschema "github.com/beto-ouverney/talker-manager-go/schema"
)

func isTokenValid(e string) bool {
	tokenRegex := regexp.MustCompile(`[\da-zA-Z]{16}`)
	return tokenRegex.MatchString(e)
}

//TokenValidate is a middleware that validates the token
func TokenValidate(header map[string][]string, decoder *json.Decoder) (ok bool, status int, message string) {
	if header["Authorization"] == nil {
		ok = false
		status = errorschema.ErrorResponse["tokenNotFound"].Status
		message = errorschema.ErrorResponse["tokenNotFound"].Message
		return
	}
	token := header["Authorization"][0]

	if len(token) == 0 {
		ok = false
		status = errorschema.ErrorResponse["tokenNotFound"].Status
		message = errorschema.ErrorResponse["tokenNotFound"].Message
		return
	} else if isTokenValid(token) == false {
		ok = false
		status = errorschema.ErrorResponse["invalidToken"].Status
		message = errorschema.ErrorResponse["invalidToken"].Message
		return
	}
	return true, 200, ""
}
