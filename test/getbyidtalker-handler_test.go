package handler

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

type Talk struct {
	WatchedAt string `json:"watchedAt"`
	Rate      int    `json:"rate"`
}

type Talker struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	ID   int    `json:"id"`
	Talk Talk   `json:"talk"`
}

type TestError struct {
	Message string `json:"message"`
}

func TestGetTalkerByIDHandler(t *testing.T) {

	router := &myrouter.Router{}
	router.Route(http.MethodGet, `/talkers/(?P<id>\d+)`, nil, handler.GetTalkerByIDHandler)

	tests := []struct {
		name            string
		describe        string
		id              string
		expectedStatus  int
		expectedMessage interface{}
	}{
		{name: "Test 1.1",
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
			t.Log(tt.describe)
			path := fmt.Sprintf("/talkers/%s", tt.id)
			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			var actual Talker
			body := json.NewDecoder(rr.Body)
			err = body.Decode(&actual)
			assert.Equal(t, tt.expectedMessage, actual)
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

		{name: "Test 2.1",
			describe:        " => Describe: It will be validated that the endpoint returns an error if no speaker is found",
			id:              "999",
			expectedStatus:  404,
			expectedMessage: TestError{"Talker not found"}},
	}
	for _, tt := range testsError {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.describe)
			path := fmt.Sprintf("/talkers/%s", tt.id)
			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			var actual TestError
			body := json.NewDecoder(rr.Body)
			err = body.Decode(&actual)
			assert.Equal(t, tt.expectedMessage, actual)
		})
	}
}
