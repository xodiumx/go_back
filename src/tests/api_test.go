package tests

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/xodiumx/go_back/api"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var router *gin.Engine

// setupRouter
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/v2/users/:id", api.GetUserHandlerWithSchema)
	router.POST("/api/v1/create", api.CreateHandler)
	return router
}

// Setup tests
func TestMain(m *testing.M) {
	router = setupRouter()
	os.Exit(m.Run())
}

func TestGetUserHandlerWithSchema(t *testing.T) {

	// Run test parallel
	time.Sleep(time.Second * 2)

	// Create client
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v2/users/1", nil)

	// Make a request
	router.ServeHTTP(w, req)

	// Check status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check body
	expectedBody := `{"id":1,"name":"John Doe","email":"john.doe@example.com"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestCreateHandler(t *testing.T) {

	// Run test parallel
	t.Parallel()
	time.Sleep(time.Second * 2)

	tests := []struct {
		name         string
		input        string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Valid JSON",
			input:        `{"name": "Alice", "age": 25}`,
			expectedCode: http.StatusOK,
			expectedBody: `{"received":{"name":"Alice","age":25}}`,
		},
		{
			name:         "Invalid JSON",
			input:        `{"name": "Alice", "age":}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: ``,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/api/v1/create", bytes.NewBuffer([]byte(tt.input)))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			if tt.expectedBody != "" {
				assert.JSONEq(t, tt.expectedBody, w.Body.String())
			}
		})
	}
}
