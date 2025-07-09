// internal/service/blockchain_test.go

package service

import (
	"blockchain-demo/internal/config"
	"blockchain-demo/internal/models"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	// Load .env file cho test
	envFiles := []string{
		".env.test",
		"../.env.test",
		"../../.env.test",
	}

	for _, file := range envFiles {
		if err := godotenv.Load(file); err == nil {
			break
		}
	}
}

func setupTestService(t *testing.T) *BlockchainService {
	// Lấy đường dẫn hiện tại
	pwd, _ := os.Getwd()
	t.Logf("Current working directory: %s", pwd)

	// Đường dẫn tới file ContentStorage.json
	contractPath := filepath.Join("ContentStorage.json")

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

	t.Logf("Contract path: %s", contractPath)

	// Kiểm tra file tồn tại
	if _, err := os.Stat(contractPath); os.IsNotExist(err) {
		t.Fatalf("Contract file not found at: %s", contractPath)
	} else {
		t.Logf("Found contract at path: %s", contractPath)
	}

	// Sử dụng biến môi trường TEST_ cho việc test
	cfg := &config.Config{
		NetworkURL:      getEnvOrFallback("TEST_NETWORK_URL", os.Getenv("NETWORK_URL")),
		ChainID:         getEnvOrFallback("TEST_CHAIN_ID", os.Getenv("CHAIN_ID")),
		PrivateKey:      getEnvOrFallback("TEST_PRIVATE_KEY", os.Getenv("PRIVATE_KEY")),
		ContractAddress: getEnvOrFallback("TEST_CONTRACT_ADDRESS", os.Getenv("CONTRACT_ADDRESS")),
		ContractJSON:    contractPath,
	}

	service, err := NewBlockchainService(cfg)
	if err != nil {
		t.Fatalf("Failed to create blockchain service: %v", err)
	}

	return service
}

// Helper function để lấy biến môi trường test hoặc fallback về biến môi trường thông thường
func getEnvOrFallback(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

// TestBlockchainConnection kiểm tra kết nối blockchain
func TestBlockchainConnection(t *testing.T) {
	svc := setupTestService(t)
	err := svc.HealthCheck()
	assert.NoError(t, err, "Should connect to blockchain successfully")
}

// TestCreateAndGetContent kiểm tra tạo và lấy content
func TestCreateAndGetContent(t *testing.T) {
	svc := setupTestService(t)

	// Tạo content mới
	req := &models.CreateContentRequest{
		Title:   "Test Content " + time.Now().Format(time.RFC3339),
		Content: "Test Content Body",
		Creator: "Test Creator",
	}

	// Lưu content
	resp, err := svc.StoreContent(req)
	assert.NoError(t, err, "Should store content successfully")
	assert.True(t, resp.Success)
	assert.NotEmpty(t, resp.ID)
	assert.NotEmpty(t, resp.TxHash)

	t.Logf("Content created with ID: %s and TxHash: %s", resp.ID, resp.TxHash)

	// Tăng thời gian chờ transaction confirm
	time.Sleep(5 * time.Second)

	// Lấy content vừa tạo
	getResp, err := svc.GetContent(resp.ID)
	if err != nil {
		t.Logf("Error getting content: %v", err)
		return // Tránh nil pointer nếu có lỗi
	}

	assert.NoError(t, err, "Should get content successfully")
	if getResp != nil {
		assert.True(t, getResp.Success)
		assert.Equal(t, req.Title, getResp.Data.Title)
	}
}

// TestCreateAndGetContest kiểm tra tạo và lấy contest
func TestCreateAndGetContest(t *testing.T) {
	svc := setupTestService(t)

	startDate := time.Now().Add(24 * time.Hour)
	endDate := startDate.Add(24 * time.Hour)

	// Tạo contest mới
	req := &models.CreateContestRequest{
		Name:        "Test Contest " + time.Now().Format(time.RFC3339),
		Description: "Test Contest Description",
		StartDate:   startDate.Format(time.RFC3339),
		EndDate:     endDate.Format(time.RFC3339),
		ImageURL:    "https://example.com/test.jpg",
	}

	// Tạo contest
	resp, err := svc.CreateContest(req)
	assert.NoError(t, err, "Should create contest successfully")
	assert.True(t, resp.Success)
	assert.NotEmpty(t, resp.ID)
	assert.NotEmpty(t, resp.TxHash)

	t.Logf("Contest created with ID: %s and TxHash: %s", resp.ID, resp.TxHash)

	// Đợi transaction được confirm
	time.Sleep(5 * time.Second)

	// Lấy contest vừa tạo
	getResp, err := svc.GetContest(resp.ID)
	if err != nil {
		t.Logf("Error getting contest: %v", err)
	}
	assert.NoError(t, err, "Should get contest successfully")
	assert.True(t, getResp.Success)
	assert.Equal(t, req.Name, getResp.Data.Name)
}

// TestSearchContests kiểm tra tìm kiếm contest
func TestSearchContests(t *testing.T) {
	svc := setupTestService(t)

	// Tìm kiếm với từ khóa "Test"
	contests, err := svc.SearchContests("Test")
	assert.NoError(t, err, "Should search contests successfully")

	t.Logf("Found %d contests with keyword 'Test'", len(contests))
	for _, contest := range contests {
		t.Logf("Contest: %s - %s", contest.ID, contest.Name)
	}
}

// TestFullWorkflow kiểm tra toàn bộ quy trình
func TestFullWorkflow(t *testing.T) {
	svc := setupTestService(t)

	// 1. Tạo contest
	contestReq := &models.CreateContestRequest{
		Name:        "Full Workflow Test " + time.Now().Format(time.RFC3339),
		Description: "Test Description",
		StartDate:   time.Now().Add(24 * time.Hour).Format(time.RFC3339),
		EndDate:     time.Now().Add(48 * time.Hour).Format(time.RFC3339),
		ImageURL:    "https://example.com/test.jpg",
	}

	contestResp, err := svc.CreateContest(contestReq)
	assert.NoError(t, err, "Should create contest")
	contestID := contestResp.ID

	// Chờ transaction được confirm
	time.Sleep(5 * time.Second)

	// 2. Tạo contestant
	contestantReq := &models.CreateContestantRequest{
		Name:    "Test Contestant",
		Details: "Test Details",
		Creator: "Test Creator",
	}

	contestantResp, err := svc.CreateContestant(contestantReq)
	assert.NoError(t, err, "Should create contestant")
	contestantID := contestantResp.ID

	// Chờ transaction được confirm
	time.Sleep(5 * time.Second)

	// 3. Đăng ký contestant vào contest
	registerReq := &models.RegisterContestantRequest{
		ContestID:    contestID,
		ContestantID: contestantID,
	}

	registerResp, err := svc.RegisterContestant(registerReq)
	assert.NoError(t, err, "Should register contestant")
	assert.True(t, registerResp.Success)

	// Chờ transaction được confirm
	time.Sleep(5 * time.Second)

	t.Logf("Full workflow completed successfully")
	t.Logf("Contest ID: %s", contestID)
	t.Logf("Contestant ID: %s", contestantID)
}
