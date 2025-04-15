package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"third-party-int/service/mock"
	inmemory "third-party-int/store/in-memory"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestStoreJokes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		Name            string
		Fail            bool
		ExpectedStatus  int
		ExpectedMessage string
	}{
		{
			"success",
			false,
			http.StatusOK,
			"Joke stored successfully",
		},
		{
			"failure",
			true,
			http.StatusInternalServerError,
			"Could not fetch jokes",
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			mockJokeService := &mock.MockJokes{Fail: tt.Fail}

			jokeStore := inmemory.NewInmemoryJokeStore()
			handler := NewJokeHandler(jokeStore, mockJokeService)
			r := gin.Default()
			r.GET("/jokes", handler.StoreJokes)
			req, _ := http.NewRequest("GET", "/jokes", nil)
			resp := httptest.NewRecorder()

			r.ServeHTTP(resp, req)
			assert.Equal(t, tt.ExpectedStatus, resp.Code)

			var body map[string]interface{}
			err := json.NewDecoder(resp.Body).Decode(&body)
			assert.NoError(t, err)

			assert.Equal(t, tt.ExpectedMessage, body["message"])
		})

	}
}
