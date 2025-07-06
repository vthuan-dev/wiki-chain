package models

import (
	"time"
)

// Content represents the data structure for general content stored on blockchain
type Content struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Creator   string    `json:"creator"`
	Timestamp time.Time `json:"timestamp"`
	TxHash    string    `json:"tx_hash,omitempty"`
	Verified  bool      `json:"verified"`
}

// Contest represents a contest/event structure
type Contest struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Organizer   string    `json:"organizer"`
	Active      bool      `json:"active"`
	ImageURL    string    `json:"image_url,omitempty"`
	TxHash      string    `json:"tx_hash,omitempty"`
	Timestamp   time.Time `json:"timestamp"`
}

// Contestant represents a participant in contests
type Contestant struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Details   string    `json:"details"`
	Creator   string    `json:"creator"`
	Timestamp time.Time `json:"timestamp"`
	Verified  bool      `json:"verified"`
	TxHash    string    `json:"tx_hash,omitempty"`
}

// Sponsor represents a sponsor for contests
type Sponsor struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	ContactInfo       string    `json:"contact_info"`
	SponsorshipAmount uint64    `json:"sponsorship_amount"`
	WalletAddress     string    `json:"wallet_address"`
	TxHash            string    `json:"tx_hash,omitempty"`
	Timestamp         time.Time `json:"timestamp"`
}

// ContestRegistration represents contestant registration for a contest
type ContestRegistration struct {
	ContestID    string    `json:"contest_id"`
	ContestantID string    `json:"contestant_id"`
	RegisteredAt time.Time `json:"registered_at"`
	TxHash       string    `json:"tx_hash,omitempty"`
}

// ============ REQUEST STRUCTS ============

// CreateContentRequest represents the request payload for creating content
type CreateContentRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Creator string `json:"creator,omitempty"`
}

// CreateContestRequest represents the request payload for creating a contest
type CreateContestRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	StartDate   string `json:"start_date" binding:"required"` // Format: "2006-01-02T15:04:05Z"
	EndDate     string `json:"end_date" binding:"required"`
	ImageURL    string `json:"image_url,omitempty"`
}

// CreateContestantRequest represents the request payload for creating a contestant
type CreateContestantRequest struct {
	Name    string `json:"name" binding:"required"`
	Details string `json:"details" binding:"required"`
	Creator string `json:"creator,omitempty"`
}

// CreateSponsorRequest represents the request payload for creating a sponsor
type CreateSponsorRequest struct {
	Name              string `json:"name" binding:"required"`
	ContactInfo       string `json:"contact_info" binding:"required"`
	SponsorshipAmount uint64 `json:"sponsorship_amount"`
}

// RegisterContestantRequest represents the request for registering contestant to contest
type RegisterContestantRequest struct {
	ContestID    string `json:"contest_id" binding:"required"`
	ContestantID string `json:"contestant_id" binding:"required"`
}

// ============ RESPONSE STRUCTS ============

// CreateContentResponse represents the response after creating content
type CreateContentResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	TxHash  string `json:"tx_hash,omitempty"`
	ID      string `json:"id,omitempty"`
}

// CreateContestResponse represents the response after creating contest
type CreateContestResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	TxHash  string `json:"tx_hash,omitempty"`
	ID      string `json:"id,omitempty"`
}

// CreateContestantResponse represents the response after creating contestant
type CreateContestantResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	TxHash  string `json:"tx_hash,omitempty"`
	ID      string `json:"id,omitempty"`
}

// CreateSponsorResponse represents the response after creating sponsor
type CreateSponsorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	TxHash  string `json:"tx_hash,omitempty"`
	ID      string `json:"id,omitempty"`
}

// RegisterContestantResponse represents the response after registering contestant
type RegisterContestantResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	TxHash  string `json:"tx_hash,omitempty"`
}

// GetContentResponse represents the response when getting content
type GetContentResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message,omitempty"`
	Data    *Content `json:"data,omitempty"`
}

// GetContestResponse represents the response when getting contest
type GetContestResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message,omitempty"`
	Data    *Contest `json:"data,omitempty"`
}

// GetContestantResponse represents the response when getting contestant
type GetContestantResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    *Contestant `json:"data,omitempty"`
}

// GetSponsorResponse represents the response when getting sponsor
type GetSponsorResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message,omitempty"`
	Data    *Sponsor `json:"data,omitempty"`
}

// ListContentsResponse represents the response when listing contents
type ListContentsResponse struct {
	Success bool       `json:"success"`
	Message string     `json:"message,omitempty"`
	Data    []*Content `json:"data,omitempty"`
	Total   int        `json:"total"`
}

// ListContestsResponse represents the response when listing contests
type ListContestsResponse struct {
	Success bool       `json:"success"`
	Message string     `json:"message,omitempty"`
	Data    []*Contest `json:"data,omitempty"`
	Total   int        `json:"total"`
}

// ListContestantsResponse represents the response when listing contestants
type ListContestantsResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message,omitempty"`
	Data    []*Contestant `json:"data,omitempty"`
	Total   int           `json:"total"`
}

// ListSponsorsResponse represents the response when listing sponsors
type ListSponsorsResponse struct {
	Success bool       `json:"success"`
	Message string     `json:"message,omitempty"`
	Data    []*Sponsor `json:"data,omitempty"`
	Total   int        `json:"total"`
}

// ListContestantsInContestResponse represents the response when getting contestants in a contest
type ListContestantsInContestResponse struct {
	Success     bool          `json:"success"`
	Message     string        `json:"message,omitempty"`
	ContestID   string        `json:"contest_id"`
	Contestants []*Contestant `json:"contestants,omitempty"`
	Total       int           `json:"total"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// ============ UTILITY STRUCTS ============

// BlockchainStats represents general statistics from blockchain
type BlockchainStats struct {
	TotalContents      int `json:"total_contents"`
	TotalContests      int `json:"total_contests"`
	TotalContestants   int `json:"total_contestants"`
	TotalSponsors      int `json:"total_sponsors"`
	TotalRegistrations int `json:"total_registrations"`
}

// BlockchainStatsResponse represents the response for blockchain statistics
type BlockchainStatsResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message,omitempty"`
	Data    *BlockchainStats `json:"data,omitempty"`
}
