package tests

import (
	"blockchain-demo/internal/config"
	"blockchain-demo/internal/models"
	"blockchain-demo/internal/service"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
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

func setupBenchmarkService(b *testing.B) *service.BlockchainService {
	// Danh sách các vị trí có thể chứa ContentStorage.json
	paths := []string{
		"ContentStorage.json",
		"../ContentStorage.json",
		"../../ContentStorage.json",
		"../real_blockchain/ContentStorage.json",
		"../../real_blockchain/ContentStorage.json",
		"../contracts/ContentStorage.json",
		"../../contracts/ContentStorage.json",
	}
	var contractPath string
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			contractPath = path
			break
		}
	}
	if contractPath == "" {
		b.Fatalf("Không tìm thấy ContentStorage.json ở bất kỳ vị trí nào! Đã thử: %v", paths)
	}
	b.Logf("DEBUG: contractPath = %s", contractPath)

	// Khởi tạo config
	cfg := &config.Config{
		NetworkURL:      os.Getenv("TEST_NETWORK_URL"),
		ChainID:         os.Getenv("TEST_CHAIN_ID"),
		PrivateKey:      os.Getenv("TEST_PRIVATE_KEY"),
		ContractAddress: os.Getenv("TEST_CONTRACT_ADDRESS"),
		ContractJSON:    contractPath,
	}

	svc, err := service.NewBlockchainService(cfg)
	if err != nil {
		b.Fatalf("Không thể khởi tạo blockchain service: %v", err)
	}

	return svc
}

// BenchmarkCreateContest đo hiệu năng tạo contest
func BenchmarkCreateContest(b *testing.B) {
	svc := setupBenchmarkService(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		contestReq := &models.CreateContestRequest{
			Name:        fmt.Sprintf("Benchmark Contest %d", i),
			Description: "Test Description for benchmark",
			StartDate:   time.Now().Add(24 * time.Hour).Format(time.RFC3339),
			EndDate:     time.Now().Add(48 * time.Hour).Format(time.RFC3339),
			ImageURL:    "https://example.com/test.jpg",
		}

		resp, err := svc.CreateContest(contestReq)
		if err != nil {
			b.Logf("Error creating contest: %v", err)
			continue
		}

		if !resp.Success {
			b.Logf("Contest creation failed")
			continue
		}

		// Đợi transaction được confirm
		time.Sleep(2 * time.Second)
	}
}

// BenchmarkSearchContests đo hiệu năng tìm kiếm contest
func BenchmarkSearchContests(b *testing.B) {
	svc := setupBenchmarkService(b)

	// Tạo một số contest trước để có dữ liệu tìm kiếm
	keywords := []string{"test", "benchmark", "contest", "performance"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		keyword := keywords[i%len(keywords)]
		contests, err := svc.SearchContests(keyword)
		if err != nil {
			b.Logf("Error searching contests: %v", err)
			continue
		}
		b.Logf("Found %d contests with keyword '%s'", len(contests), keyword)
	}
}

// BenchmarkFullWorkflow đo hiệu năng toàn bộ quy trình
func BenchmarkFullWorkflow(b *testing.B) {
	svc := setupBenchmarkService(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 1. Tạo contest
		contestReq := &models.CreateContestRequest{
			Name:        fmt.Sprintf("Benchmark Full Workflow %d", i),
			Description: "Test Description for benchmark",
			StartDate:   time.Now().Add(24 * time.Hour).Format(time.RFC3339),
			EndDate:     time.Now().Add(48 * time.Hour).Format(time.RFC3339),
			ImageURL:    "https://example.com/test.jpg",
		}

		contestResp, err := svc.CreateContest(contestReq)
		if err != nil {
			b.Logf("Error creating contest: %v", err)
			continue
		}
		time.Sleep(2 * time.Second)

		// 2. Tạo contestant
		contestantReq := &models.CreateContestantRequest{
			Name:    fmt.Sprintf("Benchmark Contestant %d", i),
			Details: "Test Details",
			Creator: "Benchmark Creator",
		}

		contestantResp, err := svc.CreateContestant(contestantReq)
		if err != nil {
			b.Logf("Error creating contestant: %v", err)
			continue
		}
		time.Sleep(2 * time.Second)

		// 3. Đăng ký contestant vào contest
		registerReq := &models.RegisterContestantRequest{
			ContestID:    contestResp.ID,
			ContestantID: contestantResp.ID,
		}

		_, err = svc.RegisterContestant(registerReq)
		if err != nil {
			b.Logf("Error registering contestant: %v", err)
			continue
		}
		time.Sleep(2 * time.Second)

		b.Logf("Full workflow completed - Contest: %s, Contestant: %s",
			contestResp.ID, contestantResp.ID)
	}
}

// BenchmarkConcurrentContestCreation đo hiệu năng tạo contest đồng thời
func BenchmarkConcurrentContestCreation(b *testing.B) {
	svc := setupBenchmarkService(b)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			contestReq := &models.CreateContestRequest{
				Name:        fmt.Sprintf("Benchmark Concurrent Contest %d", i),
				Description: "Test Description for concurrent benchmark",
				StartDate:   time.Now().Add(24 * time.Hour).Format(time.RFC3339),
				EndDate:     time.Now().Add(48 * time.Hour).Format(time.RFC3339),
				ImageURL:    "https://example.com/test.jpg",
			}

			resp, err := svc.CreateContest(contestReq)
			if err != nil {
				b.Logf("Error in concurrent contest creation: %v", err)
				continue
			}

			if !resp.Success {
				b.Logf("Concurrent contest creation failed")
				continue
			}

			time.Sleep(2 * time.Second)
			i++
		}
	})
}

// BenchmarkConcurrentSearch đo hiệu năng tìm kiếm đồng thời
func BenchmarkConcurrentSearch(b *testing.B) {
	svc := setupBenchmarkService(b)
	keywords := []string{"test", "benchmark", "contest", "performance"}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			keyword := keywords[i%len(keywords)]
			contests, err := svc.SearchContests(keyword)
			if err != nil {
				b.Logf("Error in concurrent search: %v", err)
				continue
			}
			b.Logf("Concurrent search found %d contests with keyword '%s'",
				len(contests), keyword)
			i++
		}
	})
}
