package tests

import (
	"blockchain-demo/internal/api"
	"blockchain-demo/internal/config"
	"blockchain-demo/internal/models"
	"blockchain-demo/internal/service"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupTestHandler(t *testing.T) *api.Handler {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Initialize blockchain service
	blockchainService, err := service.NewBlockchainService(cfg)
	if err != nil {
		t.Fatalf("Failed to initialize blockchain service: %v", err)
	}

	// Create handler
	handler := api.NewHandler(blockchainService)
	return handler
}

func TestHealthCheckHandler(t *testing.T) {
	handler := setupTestHandler(t)

	// Create request
	req, err := http.NewRequest("GET", "/api/v1/health", nil)
	assert.NoError(t, err)

	// Create response recorder
	rr := httptest.NewRecorder()

	// Create router and register handler
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/health", handler.HealthCheck).Methods("GET")

	// Serve request
	router.ServeHTTP(rr, req)

	// Check status code
	assert.Equal(t, http.StatusOK, rr.Code, "Health check should return 200")

	// Parse response
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")

	// Check response fields
	assert.Equal(t, "healthy", response["status"], "Status should be healthy")
	assert.Equal(t, "connected", response["blockchain"], "Blockchain should be connected")
}

func TestCreateContestHandler(t *testing.T) {
	handler := setupTestHandler(t)

	// Create test contest request
	contestReq := models.CreateContestRequest{
		Name:        "Test Contest",
		Description: "This is a test contest",
		StartDate:   "2025-07-05T00:00:00Z",
		EndDate:     "2025-08-05T00:00:00Z",
		ImageURL:    "https://example.com/test.jpg",
	}

	// Convert to JSON
	jsonData, err := json.Marshal(contestReq)
	assert.NoError(t, err, "Should be able to marshal request to JSON")

	// Create request
	req, err := http.NewRequest("POST", "/api/v1/contests", bytes.NewBuffer(jsonData))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	rr := httptest.NewRecorder()

	// Create router and register handler
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/contests", handler.CreateContest).Methods("POST")

	// Serve request
	router.ServeHTTP(rr, req)

	// Check status code
	assert.Equal(t, http.StatusCreated, rr.Code, "Create contest should return 201")

	// Parse response
	var response models.CreateContestResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")

	// Check response fields
	assert.True(t, response.Success, "Response should be successful")
	assert.NotEmpty(t, response.ID, "Contest ID should not be empty")
	assert.NotEmpty(t, response.TxHash, "Transaction hash should not be empty")
}

func TestCreateContestHandlerInvalidRequest(t *testing.T) {
	handler := setupTestHandler(t)

	// Create invalid contest request (missing required fields)
	contestReq := models.CreateContestRequest{
		Name: "", // Missing name
		// Missing other required fields
	}

	// Convert to JSON
	jsonData, err := json.Marshal(contestReq)
	assert.NoError(t, err, "Should be able to marshal request to JSON")

	// Create request
	req, err := http.NewRequest("POST", "/api/v1/contests", bytes.NewBuffer(jsonData))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	rr := httptest.NewRecorder()

	// Create router and register handler
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/contests", handler.CreateContest).Methods("POST")

	// Serve request
	router.ServeHTTP(rr, req)

	// Check status code
	assert.Equal(t, http.StatusBadRequest, rr.Code, "Invalid request should return 400")

	// Parse response
	var response models.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")

	// Check response fields
	assert.False(t, response.Success, "Response should not be successful")
	assert.Contains(t, response.Error, "Contest name is required", "Should return name required error")
}

func TestSearchContestsHandler(t *testing.T) {
	handler := setupTestHandler(t)

	// Create request with search keyword
	req, err := http.NewRequest("GET", "/api/v1/contests/search?keyword=test", nil)
	assert.NoError(t, err)

	// Create response recorder
	rr := httptest.NewRecorder()

	// Create router and register handler
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/contests/search", handler.SearchContestsHandler).Methods("GET")

	// Serve request
	router.ServeHTTP(rr, req)

	// Check status code
	assert.Equal(t, http.StatusOK, rr.Code, "Search should return 200")

	// Parse response
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")

	// Check response fields
	assert.True(t, response["success"].(bool), "Response should be successful")
	assert.NotNil(t, response["data"], "Data should not be nil")
	assert.Equal(t, float64(0), response["total"].(float64), "Total should be 0 (no contests yet)")
}

func TestSearchContestsHandlerMissingKeyword(t *testing.T) {
	handler := setupTestHandler(t)

	// Create request without search keyword
	req, err := http.NewRequest("GET", "/api/v1/contests/search", nil)
	assert.NoError(t, err)

	// Create response recorder
	rr := httptest.NewRecorder()

	// Create router and register handler
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/contests/search", handler.SearchContestsHandler).Methods("GET")

	// Serve request
	router.ServeHTTP(rr, req)

	// Check status code
	assert.Equal(t, http.StatusBadRequest, rr.Code, "Missing keyword should return 400")

	// Parse response
	var response models.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")

	// Check response fields
	assert.False(t, response.Success, "Response should not be successful")
	assert.Contains(t, response.Error, "Missing keyword", "Should return missing keyword error")
}

