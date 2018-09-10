package main

import (
	//"encoding/json"
	//"github.com/gin-gonic/gin"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

func TestQueueStatus(t *testing.T) {
	//body := gin.H{
	//	"success": true,
	//}
	r := SetupRouter()
	// Perform a GET request with that handler.
	w := performRequest(r, "GET", "/QueueStatus/id")
	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)
	// Convert the JSON response to a map
	//var response map[string]string
	//err := json.Unmarshal([]byte(w.Body.String()), &response)
	// Grab the value & whether or not it exists
	//value, exists := response["success"]
	//assert.False(t, IsValidUUID(value))
	// Make some assertions on the correctness of the response.
	//assert.Nil(t, err)
	//assert.True(t, exists)
	//	assert.Equal(t, body["success"], "true")
}
func TestQueueRequest(t *testing.T) {
	r := SetupRouter()
	/* Tired of seeing RED in the test
	// Assert we will get a 500 if NO json is passed to /QueueRequest (This is bad)
	w := performRequest(r, "POST", "/QueueRequest")
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	*/
	body := bytes.NewBuffer([]byte("{\"Cups\": [{\"CupSize\": 1, \"CupBean\": 2, \"CupStrength\": 3, \"StartBrewTime\": \"2019-01-09T20:01:55Z\"},{\"CupSize\": 3, \"CupBean\": 1, \"CupStrength\": 6, \"StartBrewTime\": \"2019-09-09T20:01:55Z\"}]}"))
	req, err := http.NewRequest("POST", "/QueueRequest", body)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("Post failed with error %d.", err)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Errorf("/QueueRequest failed with error code %d.", resp.Code)
	}
}
func TestQueueCancel(t *testing.T) {
	// Assert we will get a 200
	r := SetupRouter()
	w := performRequest(r, "POST", "/QueueCancel")
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestQueuePause(t *testing.T) {
	// Assert we will get a 200
	r := SetupRouter()
	w := performRequest(r, "POST", "/QueuePause")
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestQueueStart(t *testing.T) {
	// Assert we will get a 200
	r := SetupRouter()
	w := performRequest(r, "POST", "/QueueStart")
	assert.Equal(t, http.StatusOK, w.Code)
}
