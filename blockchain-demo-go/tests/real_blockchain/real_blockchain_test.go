package real_blockchain

import (
	"blockchain-demo/internal/config"
	"blockchain-demo/internal/models"
	"blockchain-demo/internal/service"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Load test environment variables
	err := godotenv.Load(".env.test")
	if err != nil {
		// Try to find .env.test in parent directory
		err = godotenv.Load("../.env.test")
		if err != nil {
			// If not found, try loading from project root
			err = godotenv.Load("../../.env.test")
		}
	}

	// Run tests
	code := m.Run()
	os.Exit(code)
}

func setupTestBlockchainService(t *testing.T) *service.BlockchainService {
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

	return blockchainService
}

func TestHealthCheck(t *testing.T) {
	svc := setupTestBlockchainService(t)
	err := svc.HealthCheck()
	assert.NoError(t, err, "Health check should pass without errors")
}

func TestCreateAndGetContent(t *testing.T) {
	svc := setupTestBlockchainService(t)

	// Create test content
	contentReq := &models.CreateContentRequest{
		Title:   "Test Content",
		Content: "This is test content for blockchain demo",
		Creator: "Test Creator",
	}

	// Store content
	resp, err := svc.StoreContent(contentReq)
	assert.NoError(t, err, "Storing content should not return an error")
	assert.NotEmpty(t, resp.ID, "Content ID should not be empty")

	// Get the content back
	getResp, err := svc.GetContent(resp.ID)
	assert.NoError(t, err, "Getting content should not return an error")
	assert.Equal(t, contentReq.Title, getResp.Data.Title, "Content title should match")
	assert.Equal(t, contentReq.Content, getResp.Data.Content, "Content content should match")
}

func TestCreateAndGetContest(t *testing.T) {
	svc := setupTestBlockchainService(t)

	// Create test contest
	contestReq := &models.CreateContestRequest{
		Name:        "Test Contest",
		Description: "This is a test contest",
		StartDate:   "2025-07-05T00:00:00Z",
		EndDate:     "2025-08-05T00:00:00Z",
		ImageURL:    "https://example.com/test.jpg",
	}

	// Create contest
	resp, err := svc.CreateContest(contestReq)
	assert.NoError(t, err, "Creating contest should not return an error")
	assert.NotEmpty(t, resp.ID, "Contest ID should not be empty")

	// Get the contest back
	getResp, err := svc.GetContest(resp.ID)
	assert.NoError(t, err, "Getting contest should not return an error")
	assert.Equal(t, contestReq.Name, getResp.Data.Name, "Contest name should match")
	assert.Equal(t, contestReq.Description, getResp.Data.Description, "Contest description should match")
}
