package main 

import (
	"net/http"	
	"net/http/httptest"
	"testing"
	//"encoding/json"
	"github.com/stretchr/testify/assert"
	//"github.com/gin-gonic/gin"
	"fmt"
)
func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestMain(t *testing.T){
	//body := gin.H{
	//	"Health Status": "OK",
  //}
	server := setupRoute()
	w := performRequest(server, "GET", "/healthcheck")
	fmt.Println(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)

	//var respones map[string]string
	//err := json.Unmarshal([]byte(w.Body.String()), &response)
	//value, exists := response["Health Status"]
	//assert.Nil(t, err)
	//assert.True(t, exists)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{Health Status:OK,recv_time:2022/3/28 16:49:17.178}", w.Body.String())
}