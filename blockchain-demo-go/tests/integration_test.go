package tests

import (
	"blockchain-demo/internal/config"
	"blockchain-demo/internal/models"
	"blockchain-demo/internal/service"
	"fmt"
	"testing"
	"time"

	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
)

func setupIntegrationTest(t *testing.T) *service.BlockchainService {
	// Load .env file
	_ = godotenv.Load(".env.test")

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

func TestFullContestWorkflow(t *testing.T) {
	svc := setupIntegrationTest(t)

	// Step 1: Create a contest
	contestReq := &models.CreateContestRequest{
		Name:        "Integration Test Contest",
		Description: "This is a contest for integration testing",
		StartDate:   time.Now().Add(24 * time.Hour).Format(time.RFC3339),
		EndDate:     time.Now().Add(7 * 24 * time.Hour).Format(time.RFC3339),
		ImageURL:    "https://example.com/integration-test.jpg",
	}

	contestResp, err := svc.CreateContest(contestReq)
	assert.NoError(t, err, "Creating contest should not return an error")
	assert.True(t, contestResp.Success, "Contest creation should be successful")
	assert.NotEmpty(t, contestResp.ID, "Contest ID should not be empty")
	assert.NotEmpty(t, contestResp.TxHash, "Transaction hash should not be empty")

	contestID := contestResp.ID

	// Step 2: Create a contestant
	contestantReq := &models.CreateContestantRequest{
		Name:    "Integration Test Contestant",
		Details: "This is a contestant for integration testing",
		Creator: "Integration Test Creator",
	}

	contestantResp, err := svc.CreateContestant(contestantReq)
	assert.NoError(t, err, "Creating contestant should not return an error")
	assert.True(t, contestantResp.Success, "Contestant creation should be successful")
	assert.NotEmpty(t, contestantResp.ID, "Contestant ID should not be empty")
	assert.NotEmpty(t, contestantResp.TxHash, "Transaction hash should not be empty")

	contestantID := contestantResp.ID

	// Step 3: Create a sponsor
	sponsorReq := &models.CreateSponsorRequest{
		Name:              "Integration Test Sponsor",
		ContactInfo:       "integration-test@example.com",
		SponsorshipAmount: 5000,
	}

	sponsorResp, err := svc.CreateSponsor(sponsorReq)
	assert.NoError(t, err, "Creating sponsor should not return an error")
	assert.True(t, sponsorResp.Success, "Sponsor creation should be successful")
	assert.NotEmpty(t, sponsorResp.ID, "Sponsor ID should not be empty")
	assert.NotEmpty(t, sponsorResp.TxHash, "Transaction hash should not be empty")

	// Step 4: Register contestant for contest
	regReq := &models.RegisterContestantRequest{
		ContestID:    contestID,
		ContestantID: contestantID,
	}

	regResp, err := svc.RegisterContestant(regReq)
	assert.NoError(t, err, "Registering contestant should not return an error")
	assert.True(t, regResp.Success, "Registration should be successful")
	assert.NotEmpty(t, regResp.TxHash, "Transaction hash should not be empty")

	// Step 5: Test search functionality
	searchResults, err := svc.SearchContests("Integration")
	assert.NoError(t, err, "Search should not return an error")
	assert.NotNil(t, searchResults, "Search results should not be nil")
	// Note: Currently returns empty results as GetAllContests is not fully implemented

	// Step 6: Get blockchain statistics
	statsResp, err := svc.GetBlockchainStats()
	assert.NoError(t, err, "Getting stats should not return an error")
	assert.True(t, statsResp.Success, "Stats should be successful")
	assert.NotNil(t, statsResp.Data, "Stats data should not be nil")

	// Step 7: Health check
	err = svc.HealthCheck()
	assert.NoError(t, err, "Health check should pass")

	t.Logf("✅ Integration test completed successfully")
	t.Logf("   Contest ID: %s", contestID)
	t.Logf("   Contestant ID: %s", contestantID)
	t.Logf("   Sponsor ID: %s", sponsorResp.ID)
}

func TestErrorHandling(t *testing.T) {
	svc := setupIntegrationTest(t)

	// Test invalid date format
	invalidDateReq := &models.CreateContestRequest{
		Name:        "Test Contest",
		Description: "This is a test contest",
		StartDate:   "invalid-date", // Invalid format
		EndDate:     "2025-08-05T00:00:00Z",
		ImageURL:    "https://example.com/test.jpg",
	}

	resp, err := svc.CreateContest(invalidDateReq)
	assert.Error(t, err, "Should return error for invalid date format")
	assert.False(t, resp.Success, "Should not be successful")
	assert.Contains(t, resp.Message, "Invalid start date format", "Should return date format error")

	// Test invalid date range
	invalidRangeReq := &models.CreateContestRequest{
		Name:        "Test Contest",
		Description: "This is a test contest",
		StartDate:   "2025-08-05T00:00:00Z",
		EndDate:     "2025-07-05T00:00:00Z", // End before start
		ImageURL:    "https://example.com/test.jpg",
	}

	resp, err = svc.CreateContest(invalidRangeReq)
	assert.Error(t, err, "Should return error for invalid date range")
	assert.False(t, resp.Success, "Should not be successful")
	assert.Contains(t, resp.Message, "End date must be after start date", "Should return date range error")
}

func TestSearchFunctionality(t *testing.T) {
	svc := setupIntegrationTest(t)

	// Test search with different keywords
	testCases := []string{
		"test",
		"contest",
		"blockchain",
		"integration",
		"", // Empty keyword
	}

	for _, keyword := range testCases {
		results, err := svc.SearchContests(keyword)
		assert.NoError(t, err, "Search should not return an error for keyword: %s", keyword)
		assert.NotNil(t, results, "Search results should not be nil for keyword: %s", keyword)
		// Currently returns empty results as GetAllContests is not fully implemented
	}
}

func TestConcurrentOperations(t *testing.T) {
	svc := setupIntegrationTest(t)

	// Test concurrent contest creation
	const numContests = 5
	results := make(chan *models.CreateContestResponse, numContests)
	errors := make(chan error, numContests)

	for i := 0; i < numContests; i++ {
		go func(index int) {
			contestReq := &models.CreateContestRequest{
				Name:        fmt.Sprintf("Concurrent Contest %d", index),
				Description: fmt.Sprintf("This is concurrent contest %d", index),
				StartDate:   time.Now().Add(24 * time.Hour).Format(time.RFC3339),
				EndDate:     time.Now().Add(7 * 24 * time.Hour).Format(time.RFC3339),
				ImageURL:    fmt.Sprintf("https://example.com/concurrent-%d.jpg", index),
			}

			resp, err := svc.CreateContest(contestReq)
			if err != nil {
				errors <- err
			} else {
				results <- resp
			}
		}(i)
	}

	// Collect results
	successCount := 0
	errorCount := 0

	for i := 0; i < numContests; i++ {
		select {
		case resp := <-results:
			assert.True(t, resp.Success, "Concurrent contest creation should be successful")
			assert.NotEmpty(t, resp.ID, "Contest ID should not be empty")
			successCount++
		case err := <-errors:
			t.Logf("Concurrent contest creation error: %v", err)
			errorCount++
		}
	}

	t.Logf("Concurrent test results: %d successful, %d errors", successCount, errorCount)
	assert.Greater(t, successCount, 0, "At least one concurrent operation should succeed")
}
