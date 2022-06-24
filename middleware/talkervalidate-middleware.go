package middleware

import (
	errorschema "github.com/beto-ouverney/talker-manager-go/schema"
	talker "github.com/beto-ouverney/talker-manager-go/talker/entity"
)

func checkEmptyValues(name *string, age *int, talk *talker.Talk) errorschema.StatusMsgError {
	if name == nil || len(*name) == 0 {
		return errorschema.ErrorResponse["nameIsRequires"]
	}
	if age == nil {
		return errorschema.ErrorResponse["ageIsRequired"]
	}
	if talk == nil {
		return errorschema.ErrorResponse["talkIsRequired"]
	}
	return errorschema.ErrorResponse["talkIsRequired"]
}
