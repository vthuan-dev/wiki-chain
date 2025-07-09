package service

import (
	"blockchain-demo/internal/models"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// BlockchainServiceInterface định nghĩa các phương thức cần thiết cho blockchain service
type BlockchainServiceInterface interface {
	// Content operations
	StoreContent(req *models.CreateContentRequest) (*models.CreateContentResponse, error)
	GetContent(id string) (*models.GetContentResponse, error)
	GetAllContents() (*models.ListContentsResponse, error)

	// Contest operations
	CreateContest(req *models.CreateContestRequest) (*models.CreateContestResponse, error)
	GetContest(id string) (*models.GetContestResponse, error)
	GetAllContests() (*models.ListContestsResponse, error)
	SearchContests(keyword string) ([]*models.Contest, error)

	// Contestant operations
	CreateContestant(req *models.CreateContestantRequest) (*models.CreateContestantResponse, error)
	GetContestant(id string) (*models.GetContestantResponse, error)
	GetAllContestants() (*models.ListContestantsResponse, error)

	// Sponsor operations
	CreateSponsor(req *models.CreateSponsorRequest) (*models.CreateSponsorResponse, error)
	GetSponsor(id string) (*models.GetSponsorResponse, error)
	GetAllSponsors() (*models.ListSponsorsResponse, error)

	// Registration operations
	RegisterContestant(req *models.RegisterContestantRequest) (*models.RegisterContestantResponse, error)
	GetContestantsInContest(contestID string) (*models.ListContestantsInContestResponse, error)
	IsContestantRegistered(contestID, contestantID string) (bool, error)

	// Utils
	GetBlockchainStats() (*models.BlockchainStatsResponse, error)
	HealthCheck() error
}

// MockBlockchainService là phiên bản mô phỏng của BlockchainService để dùng cho test
type MockBlockchainService struct {
	contents      map[string]*models.Content
	contests      map[string]*models.Contest
	contestants   map[string]*models.Contestant
	sponsors      map[string]*models.Sponsor
	registrations map[string]map[string]bool
}

// NewMockBlockchainService tạo instance mới của MockBlockchainService
func NewMockBlockchainService() BlockchainServiceInterface {
	return &MockBlockchainService{
		contents:      make(map[string]*models.Content),
		contests:      make(map[string]*models.Contest),
		contestants:   make(map[string]*models.Contestant),
		sponsors:      make(map[string]*models.Sponsor),
		registrations: make(map[string]map[string]bool),
	}
}

// generateID tạo ID ngẫu nhiên
func (m *MockBlockchainService) generateID() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return fmt.Sprintf("mock_%d", time.Now().Unix())
	}
	return hex.EncodeToString(bytes)
}

// generateTxHash tạo hash giao dịch giả
func (m *MockBlockchainService) generateTxHash() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "0x" + strings.Repeat("0", 64)
	}
	return "0x" + hex.EncodeToString(bytes)
}

// StoreContent giả lập lưu trữ nội dung
func (m *MockBlockchainService) StoreContent(req *models.CreateContentRequest) (*models.CreateContentResponse, error) {
	id := m.generateID()
	content := &models.Content{
		ID:        id,
		Title:     req.Title,
		Content:   req.Content,
		Creator:   req.Creator,
		Timestamp: time.Now(),
		Verified:  true,
		TxHash:    m.generateTxHash(),
	}
	m.contents[id] = content

	return &models.CreateContentResponse{
		Success: true,
		Message: "Content created successfully in mock",
		TxHash:  content.TxHash,
		ID:      id,
	}, nil
}

// GetContent giả lập lấy nội dung
func (m *MockBlockchainService) GetContent(id string) (*models.GetContentResponse, error) {
	content, exists := m.contents[id]
	if !exists {
		return &models.GetContentResponse{
			Success: false,
			Message: "Content not found in mock",
		}, fmt.Errorf("content not found")
	}

	return &models.GetContentResponse{
		Success: true,
		Data:    content,
	}, nil
}

// GetAllContents giả lập lấy tất cả nội dung
func (m *MockBlockchainService) GetAllContents() (*models.ListContentsResponse, error) {
	contents := make([]*models.Content, 0, len(m.contents))
	for _, content := range m.contents {
		contents = append(contents, content)
	}

	return &models.ListContentsResponse{
		Success: true,
		Data:    contents,
		Total:   len(contents),
	}, nil
}

// CreateContest giả lập tạo cuộc thi
func (m *MockBlockchainService) CreateContest(req *models.CreateContestRequest) (*models.CreateContestResponse, error) {
	// Validate dates
	startDate, err := time.Parse(time.RFC3339, req.StartDate)
	if err != nil {
		return &models.CreateContestResponse{
			Success: false,
			Message: "Invalid start date format. Use RFC3339 format: 2006-01-02T15:04:05Z",
		}, err
	}
	endDate, err := time.Parse(time.RFC3339, req.EndDate)
	if err != nil {
		return &models.CreateContestResponse{
			Success: false,
			Message: "Invalid end date format. Use RFC3339 format: 2006-01-02T15:04:05Z",
		}, err
	}
	if endDate.Before(startDate) {
		return &models.CreateContestResponse{
			Success: false,
			Message: "End date must be after start date",
		}, fmt.Errorf("invalid date range")
	}

	id := m.generateID()
	txHash := m.generateTxHash()

	contest := &models.Contest{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		StartDate:   startDate,
		EndDate:     endDate,
		ImageURL:    req.ImageURL,
		Organizer:   "0xMockAddress",
		TxHash:      txHash,
		Active:      true,
		Timestamp:   time.Now(),
	}

	m.contests[id] = contest

	return &models.CreateContestResponse{
		Success: true,
		Message: "Contest created successfully in mock",
		TxHash:  txHash,
		ID:      id,
	}, nil
}

// GetContest giả lập lấy thông tin cuộc thi
func (m *MockBlockchainService) GetContest(id string) (*models.GetContestResponse, error) {
	contest, exists := m.contests[id]
	if !exists {
		return &models.GetContestResponse{
			Success: false,
			Message: "Contest not found in mock",
		}, fmt.Errorf("contest not found")
	}

	return &models.GetContestResponse{
		Success: true,
		Data:    contest,
	}, nil
}

// GetAllContests giả lập lấy tất cả cuộc thi
func (m *MockBlockchainService) GetAllContests() (*models.ListContestsResponse, error) {
	contests := make([]*models.Contest, 0, len(m.contests))
	for _, contest := range m.contests {
		contests = append(contests, contest)
	}

	return &models.ListContestsResponse{
		Success: true,
		Data:    contests,
		Total:   len(contests),
	}, nil
}

// SearchContests giả lập tìm kiếm cuộc thi
func (m *MockBlockchainService) SearchContests(keyword string) ([]*models.Contest, error) {
	results := []*models.Contest{}

	// Trả về tất cả các cuộc thi (giả lập)
	for _, contest := range m.contests {
		results = append(results, contest)
	}

	return results, nil
}

// CreateContestant giả lập tạo thí sinh
func (m *MockBlockchainService) CreateContestant(req *models.CreateContestantRequest) (*models.CreateContestantResponse, error) {
	id := m.generateID()
	txHash := m.generateTxHash()

	contestant := &models.Contestant{
		ID:        id,
		Name:      req.Name,
		Details:   req.Details,
		Creator:   req.Creator,
		Timestamp: time.Now(),
		Verified:  true,
		TxHash:    txHash,
	}

	m.contestants[id] = contestant

	return &models.CreateContestantResponse{
		Success: true,
		Message: "Contestant created successfully in mock",
		TxHash:  txHash,
		ID:      id,
	}, nil
}

// GetContestant giả lập lấy thông tin thí sinh
func (m *MockBlockchainService) GetContestant(id string) (*models.GetContestantResponse, error) {
	contestant, exists := m.contestants[id]
	if !exists {
		return &models.GetContestantResponse{
			Success: false,
			Message: "Contestant not found in mock",
		}, fmt.Errorf("contestant not found")
	}

	return &models.GetContestantResponse{
		Success: true,
		Data:    contestant,
	}, nil
}

// GetAllContestants giả lập lấy tất cả thí sinh
func (m *MockBlockchainService) GetAllContestants() (*models.ListContestantsResponse, error) {
	contestants := make([]*models.Contestant, 0, len(m.contestants))
	for _, contestant := range m.contestants {
		contestants = append(contestants, contestant)
	}

	return &models.ListContestantsResponse{
		Success: true,
		Data:    contestants,
		Total:   len(contestants),
	}, nil
}

// CreateSponsor giả lập tạo nhà tài trợ
func (m *MockBlockchainService) CreateSponsor(req *models.CreateSponsorRequest) (*models.CreateSponsorResponse, error) {
	id := m.generateID()
	txHash := m.generateTxHash()

	sponsor := &models.Sponsor{
		ID:                id,
		Name:              req.Name,
		ContactInfo:       req.ContactInfo,
		SponsorshipAmount: req.SponsorshipAmount,
		WalletAddress:     "0xMockSponsorAddress",
		TxHash:            txHash,
		Timestamp:         time.Now(),
	}

	m.sponsors[id] = sponsor

	return &models.CreateSponsorResponse{
		Success: true,
		Message: "Sponsor created successfully in mock",
		TxHash:  txHash,
		ID:      id,
	}, nil
}

// GetSponsor giả lập lấy thông tin nhà tài trợ
func (m *MockBlockchainService) GetSponsor(id string) (*models.GetSponsorResponse, error) {
	sponsor, exists := m.sponsors[id]
	if !exists {
		return &models.GetSponsorResponse{
			Success: false,
			Message: "Sponsor not found in mock",
		}, fmt.Errorf("sponsor not found")
	}

	return &models.GetSponsorResponse{
		Success: true,
		Data:    sponsor,
	}, nil
}

// GetAllSponsors giả lập lấy tất cả nhà tài trợ
func (m *MockBlockchainService) GetAllSponsors() (*models.ListSponsorsResponse, error) {
	sponsors := make([]*models.Sponsor, 0, len(m.sponsors))
	for _, sponsor := range m.sponsors {
		sponsors = append(sponsors, sponsor)
	}

	return &models.ListSponsorsResponse{
		Success: true,
		Data:    sponsors,
		Total:   len(sponsors),
	}, nil
}

// RegisterContestant giả lập đăng ký thí sinh vào cuộc thi
func (m *MockBlockchainService) RegisterContestant(req *models.RegisterContestantRequest) (*models.RegisterContestantResponse, error) {
	// Kiểm tra xem cuộc thi và thí sinh có tồn tại không
	if _, exists := m.contests[req.ContestID]; !exists {
		return &models.RegisterContestantResponse{
			Success: false,
			Message: "Contest not found in mock",
		}, fmt.Errorf("contest not found")
	}

	if _, exists := m.contestants[req.ContestantID]; !exists {
		return &models.RegisterContestantResponse{
			Success: false,
			Message: "Contestant not found in mock",
		}, fmt.Errorf("contestant not found")
	}

	// Khởi tạo map cho cuộc thi nếu chưa tồn tại
	if _, exists := m.registrations[req.ContestID]; !exists {
		m.registrations[req.ContestID] = make(map[string]bool)
	}

	// Đăng ký thí sinh
	m.registrations[req.ContestID][req.ContestantID] = true
	txHash := m.generateTxHash()

	return &models.RegisterContestantResponse{
		Success: true,
		Message: "Registration successful in mock",
		TxHash:  txHash,
	}, nil
}

// GetContestantsInContest giả lập lấy danh sách thí sinh trong cuộc thi
func (m *MockBlockchainService) GetContestantsInContest(contestID string) (*models.ListContestantsInContestResponse, error) {
	if _, exists := m.contests[contestID]; !exists {
		return &models.ListContestantsInContestResponse{
			Success:   false,
			Message:   "Contest not found in mock",
			ContestID: contestID,
		}, fmt.Errorf("contest not found")
	}

	registeredContestants := m.registrations[contestID]
	contestants := make([]*models.Contestant, 0)

	for id := range registeredContestants {
		if contestant, exists := m.contestants[id]; exists {
			contestants = append(contestants, contestant)
		}
	}

	return &models.ListContestantsInContestResponse{
		Success:     true,
		ContestID:   contestID,
		Contestants: contestants,
		Total:       len(contestants),
	}, nil
}

// IsContestantRegistered giả lập kiểm tra thí sinh đã đăng ký vào cuộc thi chưa
func (m *MockBlockchainService) IsContestantRegistered(contestID, contestantID string) (bool, error) {
	if _, exists := m.registrations[contestID]; !exists {
		return false, nil
	}

	return m.registrations[contestID][contestantID], nil
}

// GetBlockchainStats giả lập lấy thống kê blockchain
func (m *MockBlockchainService) GetBlockchainStats() (*models.BlockchainStatsResponse, error) {
	return &models.BlockchainStatsResponse{
		Success: true,
		Data: &models.BlockchainStats{
			TotalContents:      len(m.contents),
			TotalContests:      len(m.contests),
			TotalContestants:   len(m.contestants),
			TotalSponsors:      len(m.sponsors),
			TotalRegistrations: 0,
		},
	}, nil
}

// HealthCheck giả lập kiểm tra kết nối
func (m *MockBlockchainService) HealthCheck() error {
	return nil
}
