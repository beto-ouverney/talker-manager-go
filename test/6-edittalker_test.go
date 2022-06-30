package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/beto-ouverney/talker-manager-go/handler"
	"github.com/beto-ouverney/talker-manager-go/middleware"
	"github.com/beto-ouverney/talker-manager-go/myrouter"
	"github.com/stretchr/testify/assert"
)

func TestEditTalker(t *testing.T) {
	seedTalkers(t)

	router := &myrouter.Router{}
	router.Route(http.MethodPut, `/talker/(?P<id>\d+)`, []myrouter.Middleware{middleware.TokenValidate, middleware.TalkerValidate}, handler.EditTalkerHandler)
	router.Route(http.MethodPost, "/login", []myrouter.Middleware{middleware.UserValidate}, handler.GetUserTokenHandler)

	tests := []struct {
		name            string
		describe        string
		ID              string
		talker          Talker
		expectedStatus  int
		expectedMessage interface{}
		jsonParams      []string
		assertMessage   string
	}{
		{
			name:     "Test 6.1",
			describe: "It will be validated that it is possible to edit a talker ",
			ID:       "1",
			talker: Talker{
				Name: "Alberto Ouverney Paz",
				Age:  30,
				Talk: Talk{
					WatchedAt: "22/10/2025",
					Rate:      5,
				},
			},
			expectedStatus: 200,
			expectedMessage: Talker{
				Name: "Alberto Ouverney Paz",
				ID:   1,
				Age:  30,
				Talk: Talk{
					WatchedAt: "22/10/2025",
					Rate:      5,
				},
			},
			assertMessage: "Talker not edited",
		},
		{
			name:     "Test 6.2",
			describe: "It not will be validated that it is possible to edit a talker without the name field ",
			ID:       "4",
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
			name:     "Test 6.3",
			describe: "It not will be validated that it is possible to edit a talker without the age field ",
			ID:       "1",
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
			name:     "Test 6.4",
			describe: "It not will be validated that it is possible to edit a talker without the talk field ",
			ID:       "5",
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
			name:     "Test 6.5",
			describe: "It not will be validated that it is possible to edit a talker without the watchedAt field ",
			ID:       "3",
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
			name:     "Test 6.6",
			describe: "It not will be validated that it is possible to edit a talker without the rate field ",
			ID:       "1",
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
			name:     "Test 6.7",
			describe: "It not will be validated that it is possible to edit a talker with a invalid age ",
			ID:       "2",
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
			name:     "Test 6.8",
			describe: "It not will be validated that it is possible to edit a talker with a invalid rate ",
			ID:       "3",
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
			name:     "Test 6.9",
			describe: "It not will be validated that it is possible to edit a talker with a invalid watchedAt ",
			ID:       "1",
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
		{
			name:     "Test 6.9",
			describe: "It not will be validated that it is possible to edit a talker with a invalid watchedAt ",
			ID:       "1",
			talker: Talker{
				Name: "Alberto Ouverney Paz",
				Age:  30,
				Talk: Talk{
					WatchedAt: "30/30/2025",
					Rate:      5,
				},
			},
			expectedStatus:  400,
			expectedMessage: TestError{Message: "The \"watchedAt\" field must have the format \"dd/mm/yyyy\""},
			jsonParams:      []string{"\"watchedAt\"", "\"dd/mm/yyyy\""},
			assertMessage:   "WatchedAt must be a valid date with \"dd/mm/yyyy\" format",
		},
		{
			name:     "Test 6.9",
			describe: "It not will be validated that it is possible to edit a talker with a invalid watchedAt ",
			ID:       "3",
			talker: Talker{
				Name: "Alberto Ouverney Paz",
				Age:  30,
				Talk: Talk{
					WatchedAt: "2006/10/10",
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
			if err != nil {
				t.Fatal(err)
			}
			path := fmt.Sprintf("/talker/%s", test.ID)
			req, err := http.NewRequest(http.MethodPut, path, bytes.NewBuffer(data))
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Authorization", token.Token)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			assert.Equal(test.expectedStatus, rr.Code, "Status code should be equal")

			if rr.Code == 200 {
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
