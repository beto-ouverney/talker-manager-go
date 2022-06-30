package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"

	"github.com/beto-ouverney/talker-manager-go/handler"
	"github.com/beto-ouverney/talker-manager-go/middleware"
	"github.com/beto-ouverney/talker-manager-go/myrouter"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type token struct {
	Token string `json:"token"`
}

func TestLoginEndPoint(t *testing.T) {

	router := &myrouter.Router{}
	router.Route(http.MethodPost, "/login", []myrouter.Middleware{middleware.UserValidate}, handler.GetUserTokenHandler)

	test := struct {
		name            string
		describe        string
		args            []User
		expectedStatus  int
		expectedMessage interface{}
	}{
		name:     "Test 3.1",
		describe: " => It will validate that the endpoint must be able to return a token with random 16 characters ",
		args: []User{
			{Email: "deferiascomigo@gmail.com",
				Password: "12345678"},
			{Email: "nhg@gmail.com",
				Password: "13344567"},
			{Email: "trybe@hotmail.com.br",
				Password: "flamengo"},
		},
		expectedStatus: 200,
	}
	tokens := make([]token, 5)
	t.Run(test.name, func(t *testing.T) {

		assert := assert.New(t)

		t.Log(test.describe)
		for _, user := range test.args {
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
			var actual token
			body := json.NewDecoder(rr.Body)
			err = body.Decode(&actual)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(16, len(actual.Token), "the token must be have 16 characters")
			tokens = append(tokens, actual)
		}
	})
	t.Run("Test 3.2", func(t *testing.T) {

		assert := assert.New(t)

		t.Log("It will validate that the endpoint must be able to return token with random 16 characters")
		sort.Slice(tokens, func(a, b int) bool { return tokens[a].Token < tokens[b].Token })

		for i, token := range tokens {
			for _, token2 := range tokens[i+1:] {
				assert.NotEqual(t, token.Token, token2.Token, "all tokens must be different")
			}
		}
	})

}
