package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beto-ouverney/talker-manager-go/myrouter"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTalkersHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	path := fmt.Sprintf("/%s", "talkers")
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := &myrouter.Router{}
	router.Route(http.MethodGet, `/talkers`, nil, GetAllTalkersHandler)
	fmt.Println(req.URL.Path)
	fmt.Println(rr)
	router.ServeHTTP(rr, req)

	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, rr.Code, http.StatusOK)
}
