package datasources

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClient_Request(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/success" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "Success")
		} else if r.URL.Path == "/error" {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Internal Server Error")
		}
	}))
	defer testServer.Close()

	client := &client{
		httpClient: http.Client{
			Timeout: 10 * time.Second,
		},
	}

	successInfo := &RequestInfo{
		HTTPMethod: http.MethodGet,
		Url:        testServer.URL + "/success",
	}
	successResponse, err := client.Request(context.Background(), successInfo)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, successResponse.StatusCode)
	assert.Equal(t, []byte("Success"), successResponse.Body)

	errorInfo := &RequestInfo{
		HTTPMethod: http.MethodGet,
		Url:        testServer.URL + "/error",
	}
	errorResponse, err := client.Request(context.Background(), errorInfo)
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, errorResponse.StatusCode)
	assert.Nil(t, errorResponse.Body)
}

func TestIsErrorStatus(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
	}
	assert.False(t, isErrorStatus(resp))

	resp.StatusCode = http.StatusInternalServerError
	assert.True(t, isErrorStatus(resp))
}
