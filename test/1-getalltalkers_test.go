package test

import (
	"encoding/json"
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
	seedTalkers(t)

	assert := assert.New(t)

	t.Log("It will be validated that the endpoint returns an array with all registered speakers")

	jsonFile, err := os.ReadFile("./talkers.json")
	var talkers []Talker
	err = json.Unmarshal(jsonFile, &talkers)
	if err != nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}
	path := fmt.Sprintf("/%s", "talker")
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := &myrouter.Router{}
	router.Route(http.MethodGet, `/talker`, nil, handler.GetAllTalkersHandler)
	router.ServeHTTP(rr, req)

	assert.Equal(rr.Code, http.StatusOK)

	var actual []Talker
	body := json.NewDecoder(rr.Body)
	err = body.Decode(&actual)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(talkers, actual)
}
