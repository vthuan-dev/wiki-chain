package service

import (
	"blockchain-demo/internal/models"
)

// IBlockchainService defines the interface for blockchain operations
type IBlockchainService interface {
	// Content operations
	StoreContent(req *models.CreateContentRequest) (*models.CreateContentResponse, error)
	GetContent(id string) (*models.GetContentResponse, error)
	GetAllContents() (*models.ListContentsResponse, error)

	// Contest operations
	CreateContest(req *models.CreateContestRequest) (*models.CreateContestResponse, error)
	GetContest(id string) (*models.GetContestResponse, error)
	GetAllContests() (*models.ListContestsResponse, error)

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

	// Statistics
	GetBlockchainStats() (*models.BlockchainStatsResponse, error)

	// Health check
	HealthCheck() error
}
