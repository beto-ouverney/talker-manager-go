package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/beto-ouverney/talker-manager-go/handler"
	"github.com/beto-ouverney/talker-manager-go/middleware"
	"github.com/beto-ouverney/talker-manager-go/myrouter"
	"github.com/stretchr/testify/assert"
)

func TestSearchTalker(t *testing.T) {

	seedTalkers(t)

	router := &myrouter.Router{}
	router.Route(http.MethodGet, `/talker/search`, []myrouter.Middleware{middleware.TokenValidate}, handler.SearchTalkersHandler)
	router.Route(http.MethodPost, "/login", []myrouter.Middleware{middleware.UserValidate}, handler.GetUserTokenHandler)

	tests := []struct {
		name            string
		describe        string
		search          string
		expectedStatus  int
		expectedMessage interface{}
		assertMessage   string
	}{
		{
			name:           "Test 8.1",
			describe:       "It will be validated that it is possible to search a talker by name",
			search:         "c",
			expectedStatus: 200,
			expectedMessage: []Talker{

				{
					Name: "Ricardo Xavier Filho",
					Age:  33,
					ID:   3,
					Talk: Talk{
						WatchedAt: "23/10/2020",
						Rate:      5,
					},
				},
				{
					Name: "Marcos Costa",
					Age:  24,
					ID:   4,
					Talk: Talk{
						WatchedAt: "23/10/2020",
						Rate:      5,
					},
				},
			},
			assertMessage: "Error searching talkers",
		},
		{
			name:            "Test 8.2",
			describe:        "It will be validated that endpoint returns an empty array when there is no talker with the given name",
			search:          "ZXDC",
			expectedStatus:  200,
			expectedMessage: []Talker{},
			assertMessage:   "Error searching talkers",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			t.Log(test.describe)

			assert := assert.New(t)

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

			req, err := http.NewRequest("GET", "/talker/search?q="+test.search, nil)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Authorization", token.Token)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(test.expectedStatus, rr.Code, "Status code not equal")

			var talkers []Talker
			err = json.NewDecoder(rr.Body).Decode(&talkers)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(test.expectedMessage, talkers, test.assertMessage)

		})
	}

	t.Run("Test 8.3", func(t *testing.T) {

		t.Log("It will be validated that the endpoint will return an array with all the speakers if param is empty")

		assert := assert.New(t)

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

		req, err := http.NewRequest("GET", "/talker/search?q=", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", token.Token)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(200, rr.Code, "Status code not equal")

		jsonFile, err := os.Open("./talkers.json")
		if err != nil {
			t.Fatal(err)
		}

		var talkersFile []Talker
		err = json.NewDecoder(jsonFile).Decode(&talkersFile)
		if err != nil {
			t.Fatal(err)
		}
		var talkers []Talker
		err = json.NewDecoder(rr.Body).Decode(&talkers)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(talkersFile, talkers, "Error searching talkers")

	})
}
