package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beto-ouverney/talker-manager-go/handler"
	"github.com/beto-ouverney/talker-manager-go/middleware"
	"github.com/beto-ouverney/talker-manager-go/myrouter"
	"github.com/stretchr/testify/assert"
)

func TestLoginValidations(t *testing.T) {

	assert := assert.New(t)

	router := &myrouter.Router{}
	router.Route(http.MethodPost, "/login", []myrouter.Middleware{middleware.UserValidate}, handler.GetUserTokenHandler)

	type arg struct {
		describe        string
		user            User
		expectedStatus  int
		expectedMessage interface{}
	}

	test := struct {
		name string
		args []arg
	}{
		name: "Test 1.1",
		args: []arg{
			{
				describe:        " => It will be validated that it is not possible to login without the email field ",
				user:            User{Password: "12345678"},
				expectedStatus:  400,
				expectedMessage: TestError{Message: "Email is required"},
			},
			{
				describe:        " => It will be validated that it is not possible to login without the password field ",
				user:            User{Email: "nhg@gmail.com"},
				expectedStatus:  400,
				expectedMessage: TestError{Message: "Password is required"},
			},
		},
	}

	t.Run(test.name, func(t *testing.T) {
		for _, tes := range test.args {
			t.Log(tes.describe)
			data, err := json.Marshal(tes.user)
			if err != nil {
				t.Fatal(err)
			}
			req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(data))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			var actual TestError
			body := json.NewDecoder(rr.Body)
			err = body.Decode(&actual)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(tes.expectedMessage, actual)
		}
	})

	t.Run("Test 1.2", func(t *testing.T) {
		t.Log(" => It will be validated that it is not possible to login with an invalid email address ")
		user := User{Email: "not an email", Password: "12345678"}
		data, err := json.Marshal(user)
		str1 := bytes.NewBuffer(data).String()
		t.Log(str1)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(data))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		var actual TestError
		body := json.NewDecoder(rr.Body)
		err = body.Decode(&actual)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(TestError{Message: "Email is not valid"}, actual)

	})

	t.Run("Test 1.3", func(t *testing.T) {
		t.Log(" => It will be validated that it is not possible to login with the password field shorter than 6 characters")
		user := User{Email: "nhg@gmail.com", Password: "12345"}
		data, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(data))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		var actual TestError
		body := json.NewDecoder(rr.Body)
		err = body.Decode(&actual)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(TestError{Message: "Password must be at least 6 characters"}, actual)
	})
}
