package main

import (
	//"encoding/json"
	//"github.com/gin-gonic/gin"
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
	w := performRequest(r, "GET", "/QueueStatus")
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
	// Assert we will get a 500 if NO json is passed to /QueueRequest (This is bad)
	r := SetupRouter()
	w := performRequest(r, "POST", "/QueueRequest")
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
