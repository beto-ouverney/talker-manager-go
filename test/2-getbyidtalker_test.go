package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beto-ouverney/talker-manager-go/handler"
	"github.com/beto-ouverney/talker-manager-go/myrouter"
	"github.com/stretchr/testify/assert"
)

func TestGetTalkerByID(t *testing.T) {
	seedTalkers(t)

	router := &myrouter.Router{}
	router.Route(http.MethodGet, `/talker/(?P<id>\d+)`, nil, handler.GetTalkerByIDHandler)

	tests := []struct {
		name            string
		describe        string
		id              string
		expectedStatus  int
		expectedMessage interface{}
	}{
		{name: "Test 2.1",
			describe:       " => Describe: It will validate that the endpoint returns a speaker person based on the route id",
			id:             "1",
			expectedStatus: 200,
			expectedMessage: Talker{
				Name: "Henrique Albuquerque",
				Age:  62,
				ID:   1,
				Talk: Talk{
					WatchedAt: "23/10/2020",
					Rate:      5,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert := assert.New(t)

			t.Log(tt.describe)
			path := fmt.Sprintf("/talker/%s", tt.id)
			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			var actual Talker
			body := json.NewDecoder(rr.Body)
			err = body.Decode(&actual)
			assert.Equal(tt.expectedStatus, rr.Code)
			assert.Equal(tt.expectedMessage, actual)
		})
	}
	t.Log("Error tests")
	testsError := []struct {
		name            string
		describe        string
		id              string
		expectedStatus  int
		expectedMessage interface{}
	}{

		{name: "Test 2.2",
			describe:        " => Describe: It will be validated that the endpoint returns an error if no speaker is found",
			id:              "999",
			expectedStatus:  404,
			expectedMessage: TestError{"Talker not found"}},
	}
	for _, tt := range testsError {
		t.Run(tt.name, func(t *testing.T) {

			assert := assert.New(t)

			t.Log(tt.describe)
			path := fmt.Sprintf("/talker/%s", tt.id)
			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			var actual TestError
			body := json.NewDecoder(rr.Body)
			err = body.Decode(&actual)
			assert.Equal(tt.expectedStatus, rr.Code)
			assert.Equal(tt.expectedMessage, actual)
		})
	}
}
