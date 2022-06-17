package middleware

import (
	"encoding/json"
	"fmt"
	"regexp"

	userentity "github.com/beto-ouverney/talker-manager-go/user/entity"
)

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func userValidations(email string, password string) (ok bool, status int, message string) {
	if email == "" {
		return false, 400, "Email is required"
	}
	if password == "" {
		return false, 400, "Password is required"
	}
	if isEmailValid(email) == false {
		return false, 400, "Email is not valid"
	}
	if len(password) < 6 {
		return false, 400, "Password must be at least 6 characters"
	}

	return true, 200, ""
}

//UserValidate is a middleware that validates the user
func UserValidate(decoder *json.Decoder) (ok bool, status int, message string) {
	fmt.Println("MIDDLEWARE")
	var user userentity.User
	err := decoder.Decode(&user)
	if err != nil {
		return false, 400, "Invalid request"
	}
	ok, status, message = userValidations(user.Email, user.Password)
	return
}
