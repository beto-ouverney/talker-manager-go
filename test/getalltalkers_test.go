package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/beto-ouverney/talker-manager-go/handler"
	"github.com/beto-ouverney/talker-manager-go/myrouter"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTalkers(t *testing.T) {
	assert := assert.New(t)
	t.Log("It will be validated that the endpoint returns an array with all registered speakers")
	jsonFile, err := os.ReadFile("./talkers.json")
	if err != nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}
	path := fmt.Sprintf("/%s", "talkers")
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := &myrouter.Router{}
	router.Route(http.MethodGet, `/talkers`, nil, handler.GetAllTalkersHandler)
	router.ServeHTTP(rr, req)

	assert.Equal(rr.Code, http.StatusOK)
	jsonFileString := string(jsonFile[:])
	assert.Equal(rr.Body.String(), jsonFileString)
}
