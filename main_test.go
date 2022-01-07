package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

type Ping struct {
	Message string
}

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	want := Ping{Message: "pong"}
	buf, _ := ioutil.ReadAll(w.Body)
	var got Ping
	json.Unmarshal(buf, &got)

	assert.Equal(t, want, got)
}
