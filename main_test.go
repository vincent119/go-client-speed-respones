package main 

import (
	"net/http"	
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T){
	server : setupRoute()
	req , _ := http.NewRequest("GET", "/healthcheck", nil) 
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)
	expectedStatus := http.StatusOK
	expectedContent := "Status: ok"
	assert.Equal(t, expectedStatus, w.Code)
  assert.Contains(t, w.Body.String(), expectedContent)
}