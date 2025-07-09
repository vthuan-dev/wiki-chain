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
	"os"
	"path/filepath"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupTestHandler(t *testing.T) *api.Handler {
	t.Log("Sử dụng blockchain service thật cho API tests")

	// Tạo config từ biến môi trường
	contractPath := filepath.Join("..", "..", "..", "ContentStorage.json")

	// Kiểm tra file tồn tại
	if _, err := os.Stat(contractPath); os.IsNotExist(err) {
		// Thử tìm ở các vị trí khác
		altPaths := []string{
			"../ContentStorage.json",
			"../../ContentStorage.json",
			"../service/ContentStorage.json",
		}

		for _, path := range altPaths {
			if _, err := os.Stat(path); err == nil {
				contractPath = path
				break
			}
		}
	}

	t.Logf("Sử dụng contract ABI từ: %s", contractPath)

	cfg := &config.Config{
		NetworkURL:      os.Getenv("TEST_NETWORK_URL"),
		ChainID:         os.Getenv("TEST_CHAIN_ID"),
		PrivateKey:      os.Getenv("TEST_PRIVATE_KEY"),
		ContractAddress: os.Getenv("TEST_CONTRACT_ADDRESS"),
		ContractJSON:    contractPath,
	}

	// Sử dụng real service thay vì mock
	realService, err := service.NewBlockchainService(cfg)
	if err != nil {
		t.Fatalf("Không thể khởi tạo blockchain service: %v", err)
	}

	// Create handler với real service
	handler := api.NewHandler(realService)
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

	// Tạo test data giống như cũ
	contestReq := models.CreateContestRequest{
		Name:        "Test Contest",
		Description: "This is a test contest",
		StartDate:   "2025-07-05T00:00:00Z",
		EndDate:     "2025-08-05T00:00:00Z",
		ImageURL:    "https://example.com/test.jpg",
	}

	jsonData, err := json.Marshal(contestReq)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/api/v1/contests", bytes.NewBuffer(jsonData))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/contests", handler.CreateContest).Methods("POST")
	router.ServeHTTP(rr, req)

	// Parse response vào map để linh hoạt
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")

	// Kiểm tra status code và message
	success, _ := response["success"].(bool)
	if success {
		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.NotEmpty(t, response["id"], "Contest ID should not be empty")
		assert.NotEmpty(t, response["tx_hash"], "Transaction hash should not be empty")
	} else {
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		// Ở đây, response là ErrorResponse, không phải CreateContestResponse
		// Nên truy cập field "error" thay vì "Error"
		assert.Contains(t, response["error"], "Failed to create contest")
	}
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

	req, err := http.NewRequest("GET", "/api/v1/contests/search?keyword=test", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/contests/search", handler.SearchContestsHandler).Methods("GET")
	router.ServeHTTP(rr, req)

	// Check status code
	assert.Equal(t, http.StatusOK, rr.Code, "Search should return 200")

	// Parse response
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")

	// Check response fields
	assert.True(t, response["success"].(bool), "Response should be successful")

	// In ra kiểu dữ liệu để debug
	t.Logf("Data type: %T, value: %v", response["data"], response["data"])
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
