package middleware

import (
	"encoding/json"
	"regexp"

	errorschema "github.com/beto-ouverney/talker-manager-go/schema"
	talker "github.com/beto-ouverney/talker-manager-go/talker/entity"
)

func checkEmptyValues(name *string, age *int, talk *talker.Talk) (errorschema.StatusMsgError, bool) {
	if name == nil || len(*name) == 0 {
		return errorschema.ErrorResponse["nameIsRequired"], true
	}
	if age == nil {
		return errorschema.ErrorResponse["ageIsRequired"], true
	}
	if talk == nil || len(talk.WatchedAt) == 0 && talk.Rate == 0 {
		return errorschema.ErrorResponse["talkIsRequired"], true
	}
	return errorschema.StatusMsgError{}, false
}

func isDateValid(date string) bool {
	dateRegex := regexp.MustCompile(`(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\d\d)`)
	return dateRegex.MatchString(date)
}

// TalkerValidate is a middleware that validates the talker
func TalkerValidate(header map[string][]string, decoder *json.Decoder) (ok bool, status int, message string) {
	var talker *talker.Talker
	err := decoder.Decode(&talker)
	if err != nil {
		return false, 400, "Invalid request"
	}
	errorEmptyValues, erro := checkEmptyValues(&talker.Name, &talker.Age, &talker.Talk)
	if erro {
		ok = false
		status = errorEmptyValues.Status
		message = errorEmptyValues.Message
		return
	}
	if len(talker.Name) < 3 {
		ok = false
		status = errorschema.ErrorResponse["invalidName"].Status
		message = errorschema.ErrorResponse["invalidName"].Message
		return
	}

	if talker.Age < 18 {
		ok = false
		status = errorschema.ErrorResponse["invalidAge"].Status
		message = errorschema.ErrorResponse["invalidAge"].Message
		return
	}

	if isDateValid(talker.Talk.WatchedAt) == false {
		ok = false
		status = errorschema.ErrorResponse["invalidWathedAt"].Status
		message = errorschema.ErrorResponse["invalidWathedAt"].Message
		return
	}
	if talker.Talk.Rate < 1 || talker.Talk.Rate > 5 {
		ok = false
		status = errorschema.ErrorResponse["invalidRate"].Status
		message = errorschema.ErrorResponse["invalidRate"].Message
		return
	}
	return true, 200, ""
}
