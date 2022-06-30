package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/beto-ouverney/talker-manager-go/handler"
	"github.com/beto-ouverney/talker-manager-go/middleware"
	"github.com/beto-ouverney/talker-manager-go/myrouter"
	"github.com/stretchr/testify/assert"
)

func TestCreateTalker(t *testing.T) {
	seedTalkers(t)

	router := &myrouter.Router{}
	router.Route(http.MethodPost, "/talker", []myrouter.Middleware{middleware.TokenValidate, middleware.TalkerValidate}, handler.AddTalkerHandler)
	router.Route(http.MethodPost, "/login", []myrouter.Middleware{middleware.UserValidate}, handler.GetUserTokenHandler)

	type test struct {
		name            string
		describe        string
		talker          Talker
		expectedStatus  int
		expectedMessage interface{}
		jsonParams      []string
		assertMessage   string
	}

	tests := []struct {
		name            string
		describe        string
		talker          Talker
		expectedStatus  int
		expectedMessage interface{}
		jsonParams      []string
		assertMessage   string
	}{
		{
			name:     "Test 5.1",
			describe: "It will be validated that it is possible to add a talker ",
			talker: Talker{
				Name: "Alberto Ouverney Paz",
				Age:  30,
				Talk: Talk{
					WatchedAt: "22/10/2025",
					Rate:      5,
				},
			},
			expectedStatus: 201,
			expectedMessage: Talker{
				Name: "Alberto Ouverney Paz",
				ID:   5,
				Age:  30,
				Talk: Talk{
					WatchedAt: "22/10/2025",
					Rate:      5,
				},
			},
			assertMessage: "Talker not added",
		},
		{
			name:     "Test 5.2",
			describe: "It not will be validated that it is possible to add a talker without the name field ",
			talker: Talker{
				Age: 30,
				Talk: Talk{
					WatchedAt: "22/10/2025",
					Rate:      5,
				},
			},
			expectedStatus:  400,
			expectedMessage: TestError{Message: `The "name" field is required`},
			jsonParams:      []string{"\"name\""},
			assertMessage:   "Name is required",
		},
		{
			name:     "Test 5.3",
			describe: "It not will be validated that it is possible to add a talker without the age field ",
			talker: Talker{
				Name: "Alberto Ouverney Paz",
				Talk: Talk{
					WatchedAt: "22/10/2025",
					Rate:      5,
				},
			},
			expectedStatus:  400,
			expectedMessage: TestError{Message: "The speaker must be of legal age"},
			jsonParams:      []string{"\"age\""},
			assertMessage:   "Age is required",
		},
		{
			name:     "Test 5.4",
			describe: "It not will be validated that it is possible to add a talker without the talk field ",
			talker: Talker{
				Name: "Beto Ouverney Paz",
				Age:  30,
			},
			expectedStatus:  400,
			expectedMessage: TestError{Message: "The \"talk\" field is mandatory and \"watchedAt\" and \"rate\" cannot be empty"},
			jsonParams:      []string{"\"talk\"", "\"rate\"", "\"watchedAt\""},
			assertMessage:   "Talk is required",
		},
		{
			name:     "Test 5.5",
			describe: "It not will be validated that it is possible to add a talker without the watchedAt field ",
			talker: Talker{
				Name: "Beto Ouverney",
				Age:  30,
				Talk: Talk{
					Rate: 5,
				},
			},
			expectedStatus:  400,
			expectedMessage: TestError{Message: "The \"watchedAt\" field must have the format \"dd/mm/yyyy\""},
			jsonParams:      []string{"\"talk\"", "\"rate\"", "\"watchedAt\"", "\"dd/mm/yyyy\""},
			assertMessage:   "WatchedAt is required",
		},
		{
			name:     "Test 5.6",
			describe: "It not will be validated that it is possible to add a talker without the rate field ",
			talker: Talker{
				Name: "Alberto Ouverney Paz",
				Age:  30,
				Talk: Talk{
					WatchedAt: "22/10/2025",
				},
			},
			expectedStatus:  400,
			expectedMessage: TestError{Message: `The "rate" field must be an integer from 1 to 5`},
			jsonParams:      []string{"\"rate\""},
			assertMessage:   "Rate is required",
		},
		{
			name:     "Test 5.7",
			describe: "It not will be validated that it is possible to add a talker with a invalid age ",
			talker: Talker{
				Name: "Alberto Ouverney Paz",
				Age:  1,
				Talk: Talk{
					WatchedAt: "22/10/2025",
					Rate:      5,
				},
			},
			expectedStatus:  400,
			expectedMessage: TestError{Message: `The speaker must be of legal age`},
			jsonParams:      []string{"\"age\""},
			assertMessage:   "Age must be greater than or equal to 18",
		},
		{
			name:     "Test 5.8",
			describe: "It not will be validated that it is possible to add a talker with a invalid rate ",
			talker: Talker{
				Name: "Alberto Ouverney Paz",
				Age:  30,
				Talk: Talk{
					WatchedAt: "22/10/2025",
					Rate:      7,
				},
			},
			expectedStatus:  400,
			expectedMessage: TestError{Message: "The \"rate\" field must be an integer from 1 to 5"},
			jsonParams:      []string{"\"rate\""},
			assertMessage:   "The \"rate\" field must be an integer from 1 to 5",
		},
		{
			name:     "Test 5.9",
			describe: "It not will be validated that it is possible to add a talker with a invalid watchedAt ",
			talker: Talker{
				Name: "Alberto Ouverney Paz",
				Age:  30,
				Talk: Talk{
					WatchedAt: "//2025",
					Rate:      5,
				},
			},
			expectedStatus:  400,
			expectedMessage: TestError{Message: "The \"watchedAt\" field must have the format \"dd/mm/yyyy\""},
			jsonParams:      []string{"\"watchedAt\"", "\"dd/mm/yyyy\""},
			assertMessage:   "WatchedAt must be a valid date with \"dd/mm/yyyy\" format",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
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

			data, err := json.Marshal(test.talker)

			req, err := http.NewRequest("POST", "/talker", bytes.NewBuffer(data))
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Authorization", token.Token)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(test.expectedStatus, rr.Code, "Status code should be equal")

			if rr.Code == 201 {
				var actual Talker
				err = json.NewDecoder(rr.Body).Decode(&actual)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(test.expectedMessage, actual, test.assertMessage)
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
}
