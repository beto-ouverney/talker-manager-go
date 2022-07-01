package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/beto-ouverney/talker-manager-go/handler"
	"github.com/beto-ouverney/talker-manager-go/middleware"
	"github.com/beto-ouverney/talker-manager-go/myrouter"
	"github.com/stretchr/testify/assert"
)

func TestDeleteTalker(t *testing.T) {

	router := &myrouter.Router{}
	router.Route(http.MethodPost, "/talker", []myrouter.Middleware{middleware.TokenValidate, middleware.TalkerValidate}, handler.AddTalkerHandler)
	router.Route(http.MethodDelete, `/talker/(?P<id>\d+)`, []myrouter.Middleware{middleware.TokenValidate}, handler.DeleteTalkerHandler)
	router.Route(http.MethodPost, "/login", []myrouter.Middleware{middleware.UserValidate}, handler.GetUserTokenHandler)

	tests := []struct {
		name            string
		describe        string
		ID              string
		expectedStatus  int
		expectedMessage interface{}
		jsonParams      []string
		assertMessage   string
	}{
		{
			name:            "Test 7.1",
			describe:        "It will be validated that it is possible to delete a talker ",
			ID:              "5",
			expectedStatus:  204,
			expectedMessage: nil,
			assertMessage:   "Talker not deleted",
		},
		{
			name:            "Test 7.2",
			describe:        "It not will be validated that it is possible to delete a talker with an invalid id",
			ID:              "777",
			expectedStatus:  404,
			expectedMessage: TestError{Message: "Talker not found"},
			assertMessage:   "Talker not deleted",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			seedTalkers(t)
			assert := assert.New(t)
			t.Log(test.describe)

			user := User{
				Email:    "deferiascomigo@gmail.com",
				Password: "12345678",
			}
			userData, err := json.Marshal(user)
			if err != nil {
				t.Fatal(err)
			}

			reqToken, err := http.NewRequest("POST", "/login", bytes.NewBuffer(userData))
			if err != nil {
				t.Fatal(err)
			}
			rrToken := httptest.NewRecorder()
			router.ServeHTTP(rrToken, reqToken)
			var token token
			err = json.NewDecoder(rrToken.Body).Decode(&token)
			if err != nil {
				t.Fatal(err)
			}

			talkerAdd := Talker{
				Name: "Alberto Ouverney Paz",
				Age:  30,
				Talk: Talk{
					WatchedAt: "22/10/2025",
					Rate:      5,
				},
			}

			talkerAddData, err := json.Marshal(talkerAdd)
			if err != nil {
				t.Fatal(err)
			}
			reqTalkerAdd, err := http.NewRequest("POST", "/talkers", bytes.NewBuffer(talkerAddData))

			if err != nil {
				t.Fatal(err)
			}

			reqTalkerAdd.Header.Set("Authorization", token.Token)
			rrTalkerAdd := httptest.NewRecorder()
			router.ServeHTTP(rrTalkerAdd, reqTalkerAdd)

			if rrTalkerAdd.Code != 201 {
				t.Error("Error adding talker")
			}

			path := fmt.Sprintf("/talker/%s", test.ID)
			req, err := http.NewRequest(http.MethodDelete, path, nil)
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Authorization", token.Token)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			assert.Equal(test.expectedStatus, rr.Code, "Status code should be equal")

			if rr.Code == 204 {
				jsonFile, err := os.Open("./talkers.json")
				if err != nil {
					t.Fatal(err)
				}
				var talkers []Talker
				err = json.NewDecoder(jsonFile).Decode(&talkers)
				if err != nil {
					t.Fatal(err)
				}

				assert.NotContains(
					talkers, talkerAdd, test.assertMessage)
			} else {
				body := rr.Body.String()
				for _, param := range test.jsonParams {
					newWord := strings.ReplaceAll(param, `"`, `\"`)
					body = strings.ReplaceAll(body, param, newWord)
				}
				var actualErr TestError

				err = json.Unmarshal([]byte(body), &actualErr)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(test.expectedMessage, actualErr, test.assertMessage)
			}

		})
	}

	newTest := struct {
		name            string
		describe        string
		ID              string
		expectedStatus  int
		expectedMessage interface{}
		assertMessage   string
	}{

		name:            "Test 7.3",
		describe:        "It not will be validated that it is possible to delete a talker without autorization",
		ID:              "1",
		expectedStatus:  401,
		expectedMessage: TestError{Message: "Token not found"},
		assertMessage:   "Talker not deleted",
	}

	t.Run(newTest.name, func(t *testing.T) {
		seedTalkers(t)
		assert := assert.New(t)
		t.Log(newTest.describe)

		path := fmt.Sprintf("/talker/%s", newTest.ID)
		req, err := http.NewRequest(http.MethodDelete, path, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(newTest.expectedStatus, rr.Code, "Status code should be equal")
		var actual TestError
		err = json.Unmarshal(rr.Body.Bytes(), &actual)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(newTest.expectedMessage, actual, newTest.assertMessage)

	})

	t.Run("Test 7.4", func(t *testing.T) {
		t.Log("It not will be validated that it is possible to delete a talker with invalid token")

		seedTalkers(t)
		assert := assert.New(t)

		jsonFile, err := os.Open("./talkers.json")
		if err != nil {
			t.Fatal(err)
		}
		var talkersOriginal []Talker
		err = json.NewDecoder(jsonFile).Decode(&talkersOriginal)
		if err != nil {
			t.Fatal(err)
		}

		path := fmt.Sprintf("/talker/%s", newTest.ID)
		req, err := http.NewRequest(http.MethodDelete, path, nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "99999999")
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(newTest.expectedStatus, rr.Code, "Status code should be equal")
		var actual TestError
		err = json.Unmarshal(rr.Body.Bytes(), &actual)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(TestError{Message: "Token invalid"}, actual, newTest.assertMessage)

	})

}
