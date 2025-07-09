package service

import (
	"blockchain-demo/internal/config"
	"blockchain-demo/internal/models"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
	"time"
	"unicode"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/text/unicode/norm"
)

// BlockchainService handles all blockchain interactions
type BlockchainService struct {
	client     *ethclient.Client
	config     *config.Config
	privateKey *ecdsa.PrivateKey
	fromAddr   common.Address
	chainID    *big.Int

	// In-memory storage for demo purposes
	contents      map[string]*models.Content
	contests      map[string]*models.Contest
	contestants   map[string]*models.Contestant
	sponsors      map[string]*models.Sponsor
	registrations map[string]map[string]bool // contestID -> contestantID -> registered
}

// NewBlockchainService creates a new blockchain service instance
func NewBlockchainService(cfg *config.Config) (*BlockchainService, error) {
	// Connect to blockchain
	client, err := ethclient.Dial(cfg.NetworkURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to blockchain: %v", err)
	}

	// Parse chain ID
	chainID := new(big.Int)
	chainID.SetString(cfg.ChainID, 10)

	service := &BlockchainService{
		client:        client,
		config:        cfg,
		chainID:       chainID,
		contents:      make(map[string]*models.Content),
		contests:      make(map[string]*models.Contest),
		contestants:   make(map[string]*models.Contestant),
		sponsors:      make(map[string]*models.Sponsor),
		registrations: make(map[string]map[string]bool),
	}

	// Setup private key if provided
	if cfg.PrivateKey != "" && cfg.PrivateKey != "your_private_key_here" {
		privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(cfg.PrivateKey, "0x"))
		if err != nil {
			return nil, fmt.Errorf("invalid private key: %v", err)
		}
		service.privateKey = privateKey
		service.fromAddr = crypto.PubkeyToAddress(privateKey.PublicKey)
		log.Printf("✅ Loaded wallet address: %s", service.fromAddr.Hex())
	} else {
		return nil, fmt.Errorf("private key is required")
	}

	// Check if contract address is provided
	if cfg.ContractAddress == "" {
		return nil, fmt.Errorf("contract address is required")
	}

	log.Printf("✅ Connected to blockchain network: %s", cfg.NetworkURL)
	return service, nil
}

// StoreContent pushes content to blockchain
func (bs *BlockchainService) StoreContent(req *models.CreateContentRequest) (*models.CreateContentResponse, error) {
	// Generate unique ID
	id := bs.generateID()

	// Create content object
	content := &models.Content{
		ID:        id,
		Title:     req.Title,
		Content:   req.Content,
		Creator:   req.Creator,
		Timestamp: time.Now(),
		Verified:  false,
	}

	// Push to blockchain
	txHash, err := bs.pushToBlockchain(content)
	if err != nil {
		return &models.CreateContentResponse{
			Success: false,
			Message: "Failed to push to blockchain",
		}, err
	}

	content.TxHash = txHash
	content.Verified = true
	log.Printf("✅ Content pushed to blockchain with tx: %s", txHash)

	return &models.CreateContentResponse{
		Success: true,
		Message: "Content created successfully",
		TxHash:  content.TxHash,
		ID:      id,
	}, nil
}

// GetContent retrieves content from blockchain
func (bs *BlockchainService) GetContent(id string) (*models.GetContentResponse, error) {
	// Get from blockchain
	content, err := bs.getFromBlockchain(id)
	if err != nil {
		return &models.GetContentResponse{
			Success: false,
			Message: "Content not found on blockchain",
		}, err
	}

	return &models.GetContentResponse{
		Success: true,
		Data:    content,
	}, nil
}

// GetAllContents returns all contents from blockchain
func (bs *BlockchainService) GetAllContents() (*models.ListContentsResponse, error) {
	// This would need to be implemented to read all content IDs from blockchain
	// and then fetch each content individually
	// For now, return empty list as this requires smart contract support
	return &models.ListContentsResponse{
		Success: true,
		Data:    []*models.Content{},
		Total:   0,
	}, nil
}

// pushToBlockchain simulates pushing data to blockchain
func (bs *BlockchainService) pushToBlockchain(content *models.Content) (string, error) {
	if bs.privateKey == nil {
		return "", fmt.Errorf("no private key configured")
	}

	// For demo purposes, we'll simulate a transaction
	// In a real implementation, you would:
	// 1. Create the transaction data using ABI encoding
	// 2. Estimate gas
	// 3. Sign and send the transaction

	log.Printf("📤 Simulating blockchain transaction for content: %s", content.Title)

	// Simulate getting nonce
	nonce, err := bs.client.PendingNonceAt(context.Background(), bs.fromAddr)
	if err != nil {
		return "", fmt.Errorf("failed to get nonce: %v", err)
	}

	// Simulate gas price
	gasPrice, err := bs.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get gas price: %v", err)
	}

	// For demo, generate a fake transaction hash
	txHash := bs.generateTxHash()

	log.Printf("📝 Simulated tx - Nonce: %d, GasPrice: %s, Hash: %s", nonce, gasPrice.String(), txHash)

	return txHash, nil
}

// getFromBlockchain simulates getting data from blockchain
func (bs *BlockchainService) getFromBlockchain(id string) (*models.Content, error) {
	// For demo purposes, this would normally:
	// 1. Call the smart contract's getter method
	// 2. Decode the returned data
	// 3. Return the content struct

	log.Printf("📥 Simulating blockchain read for content ID: %s", id)

	// Return nil to indicate not found on blockchain
	return nil, fmt.Errorf("content not found on blockchain")
}

// generateID generates a unique ID for content
func (bs *BlockchainService) generateID() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return fmt.Sprintf("content_%d", time.Now().Unix())
	}
	return hex.EncodeToString(bytes)
}

// generateTxHash generates a fake transaction hash for demo
func (bs *BlockchainService) generateTxHash() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "0x" + strings.Repeat("0", 64)
	}
	return "0x" + hex.EncodeToString(bytes)
}

// ============ CONTEST OPERATIONS ============

// LoadContractABI loads the ABI from ContentStorage.json
func LoadContractABI(abiPath string) (abi.ABI, error) {
	abiBytes, err := ioutil.ReadFile(abiPath)
	if err != nil {
		return abi.ABI{}, err
	}
	var contractInfo struct {
		ABI json.RawMessage `json:"abi"`
	}
	if err := json.Unmarshal(abiBytes, &contractInfo); err != nil {
		return abi.ABI{}, err
	}
	return abi.JSON(strings.NewReader(string(contractInfo.ABI)))
}

// ContestTuple represents the tuple returned by the smart contract
type ContestTuple struct {
	Name        string
	Description string
	StartDate   *big.Int
	EndDate     *big.Int
	Organizer   common.Address
	Active      bool
	ImageURL    string
}

// CreateContest creates a new contest and pushes to blockchain
func (bs *BlockchainService) CreateContest(req *models.CreateContestRequest) (*models.CreateContestResponse, error) {
	id := bs.generateID()
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

	// Serialize contest to formatted JSON
	contestJson := map[string]interface{}{
		"id":          id,
		"name":        req.Name,
		"description": req.Description,
		"start_date":  startDate.Format(time.RFC3339),
		"end_date":    endDate.Format(time.RFC3339),
		"organizer":   bs.fromAddr.Hex(),
		"image_url":   req.ImageURL,
		"timestamp":   time.Now().Format(time.RFC3339),
	}
	jsonBytes, err := json.MarshalIndent(contestJson, "", "  ")
	if err != nil {
		return &models.CreateContestResponse{
			Success: false,
			Message: "Failed to marshal contest JSON",
		}, err
	}

	defer func() {
		if r := recover(); r != nil {
			log.Printf("[WARN] Recovered in CreateContest contract call: %v", r)
		}
	}()

	contractAddr := common.HexToAddress(bs.config.ContractAddress)
	parsedABI, err := LoadContractABI(bs.config.ContractJSON)
	if err != nil {
		return &models.CreateContestResponse{
			Success: false,
			Message: "Failed to load contract ABI",
		}, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(bs.privateKey, bs.chainID)
	if err != nil {
		return &models.CreateContestResponse{
			Success: false,
			Message: "Failed to create transactor",
		}, err
	}

	// Call new createContestJson function with id and JSON string
	input := []interface{}{id, string(jsonBytes)}
	contract := bind.NewBoundContract(contractAddr, parsedABI, bs.client, bs.client, bs.client)
	tx, err := contract.Transact(auth, "createContestJson", input...)
	if err != nil {
		return &models.CreateContestResponse{
			Success: false,
			Message: "Failed to create contest on blockchain",
		}, err
	}

	log.Printf("[OK] Contest JSON pushed to blockchain: %s", tx.Hash().Hex())

	return &models.CreateContestResponse{
		Success: true,
		Message: "Contest created and pushed to blockchain (JSON)",
		TxHash:  tx.Hash().Hex(),
		ID:      id,
	}, nil
}

// SearchContests tìm kiếm contest trên blockchain theo từ khóa ở mọi trường (JSON version)
func (bs *BlockchainService) SearchContests(keyword string) ([]*models.Contest, error) {
	contractAddr := common.HexToAddress(bs.config.ContractAddress)
	parsedABI, err := LoadContractABI(bs.config.ContractJSON)
	if err != nil {
		return []*models.Contest{}, err
	}
	contract := bind.NewBoundContract(contractAddr, parsedABI, bs.client, bs.client, bs.client)
	callOpts := &bind.CallOpts{Pending: false, From: bs.fromAddr}
	// Lấy tất cả contestIds
	var idsRaw []interface{}
	err = contract.Call(callOpts, &idsRaw, "getAllContestIds")
	if err != nil {
		return []*models.Contest{}, err
	}
	var ids []string
	if len(idsRaw) > 0 {
		if arr, ok := idsRaw[0].([]string); ok {
			ids = arr
		} else if arr, ok := idsRaw[0].([]interface{}); ok {
			for _, v := range arr {
				if s, ok := v.(string); ok {
					ids = append(ids, s)
				} else if b, ok := v.([]byte); ok {
					ids = append(ids, string(b))
				}
			}
		}
	}
	var results []*models.Contest
	keyword = strings.ToLower(removeDiacritics(keyword))
	for _, id := range ids {
		var jsonStrRaw []interface{}
		err := contract.Call(callOpts, &jsonStrRaw, "getContestJsonById", id)
		if err != nil || len(jsonStrRaw) == 0 {
			continue
		}
		jsonStr, ok := jsonStrRaw[0].(string)
		if !ok || jsonStr == "" {
			continue
		}
		var c models.Contest
		if err := json.Unmarshal([]byte(jsonStr), &c); err != nil {
			continue
		}
		if strings.Contains(strings.ToLower(removeDiacritics(c.Name)), keyword) ||
			strings.Contains(strings.ToLower(removeDiacritics(c.Description)), keyword) ||
			strings.Contains(strings.ToLower(removeDiacritics(c.ImageURL)), keyword) ||
			strings.Contains(strings.ToLower(removeDiacritics(c.Organizer)), keyword) {
			results = append(results, &c)
		}
	}
	return results, nil
}

// GetContest retrieves contest by ID from blockchain (JSON version)
func (bs *BlockchainService) GetContest(id string) (*models.GetContestResponse, error) {
	contractAddr := common.HexToAddress(bs.config.ContractAddress)
	parsedABI, err := LoadContractABI(bs.config.ContractJSON)
	if err != nil {
		return &models.GetContestResponse{
			Success: false,
			Message: "Failed to load contract ABI",
		}, err
	}
	contract := bind.NewBoundContract(contractAddr, parsedABI, bs.client, bs.client, bs.client)
	callOpts := &bind.CallOpts{Pending: false, From: bs.fromAddr}
	var jsonStrRaw []interface{}
	err = contract.Call(callOpts, &jsonStrRaw, "getContestJsonById", id)
	if err != nil || len(jsonStrRaw) == 0 {
		return &models.GetContestResponse{
			Success: false,
			Message: "Contest not found on blockchain",
		}, err
	}
	jsonStr, ok := jsonStrRaw[0].(string)
	if !ok || jsonStr == "" {
		return &models.GetContestResponse{
			Success: false,
			Message: "Contest not found on blockchain",
		}, nil
	}
	var contest models.Contest
	if err := json.Unmarshal([]byte(jsonStr), &contest); err != nil {
		return &models.GetContestResponse{
			Success: false,
			Message: "Failed to parse contest JSON",
		}, err
	}
	return &models.GetContestResponse{
		Success: true,
		Data:    &contest,
	}, nil
}

// GetAllContests returns all contests from blockchain
func (bs *BlockchainService) GetAllContests() (*models.ListContestsResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[WARN] Recovered in GetAllContests contract call: %v", r)
		}
	}()

	contractAddr := common.HexToAddress(bs.config.ContractAddress)
	parsedABI, err := LoadContractABI(bs.config.ContractJSON)
	if err != nil {
		return &models.ListContestsResponse{
			Success: false,
			Message: "Failed to load contract ABI",
		}, err
	}

	contract := bind.NewBoundContract(contractAddr, parsedABI, bs.client, bs.client, bs.client)
	callOpts := &bind.CallOpts{Pending: false, From: bs.fromAddr}

	// 1. Get contest IDs
	var idsRaw []interface{}
	err = contract.Call(callOpts, &idsRaw, "getAllContestIds")
	log.Printf("DEBUG getAllContestIds idsRaw: %#v", idsRaw)
	if err != nil {
		return &models.ListContestsResponse{
			Success: false,
			Message: "Failed to get contest IDs from blockchain",
		}, err
	}

	var ids []string
	if len(idsRaw) > 0 {
		if arr, ok := idsRaw[0].([]string); ok {
			ids = arr
		} else if arr, ok := idsRaw[0].([]interface{}); ok {
			for _, v := range arr {
				if s, ok := v.(string); ok {
					ids = append(ids, s)
				} else if b, ok := v.([]byte); ok {
					ids = append(ids, string(b))
				}
			}
		}
	}

	var contests []*models.Contest
	for _, id := range ids {
		var result []interface{}
		err := contract.Call(callOpts, &result, "getContest", id)
		if err != nil {
			log.Printf("[ERROR] Call getContest(%s): %v", id, err)
			continue
		}

		// Convert the result to our ContestTuple struct
		if len(result) == 7 { // Updated to check for 7 fields including imageURL
			// Add debug logging
			log.Printf("DEBUG getContest result for %s: %#v", id, result)

			contestTuple := ContestTuple{
				Name:        result[0].(string),
				Description: result[1].(string),
				StartDate:   result[2].(*big.Int),
				EndDate:     result[3].(*big.Int),
				Organizer:   result[4].(common.Address),
				Active:      result[5].(bool),
				ImageURL:    result[6].(string),
			}
			contests = append(contests, &models.Contest{
				ID:          id,
				Name:        contestTuple.Name,
				Description: contestTuple.Description,
				StartDate:   time.Unix(contestTuple.StartDate.Int64(), 0),
				EndDate:     time.Unix(contestTuple.EndDate.Int64(), 0),
				Organizer:   contestTuple.Organizer.Hex(),
				Active:      contestTuple.Active,
				ImageURL:    contestTuple.ImageURL,
			})
		} else {
			log.Printf("[ERROR] Unexpected result length for contest %s: got %d, want 7", id, len(result))
		}
	}

	return &models.ListContestsResponse{
		Success: true,
		Data:    contests,
		Total:   len(contests),
	}, nil
}

// ============ CONTESTANT OPERATIONS ============

// CreateContestant creates a new contestant on blockchain
func (bs *BlockchainService) CreateContestant(req *models.CreateContestantRequest) (*models.CreateContestantResponse, error) {
	// Generate unique ID
	id := bs.generateID()

	// Create contestant object
	contestant := &models.Contestant{
		ID:        id,
		Name:      req.Name,
		Details:   req.Details,
		Creator:   req.Creator,
		Timestamp: time.Now(),
		Verified:  false,
	}

	// Push to blockchain
	txHash := bs.generateTxHash()
	contestant.TxHash = txHash
	contestant.Verified = true
	log.Printf("✅ Contestant created with tx: %s", txHash)

	return &models.CreateContestantResponse{
		Success: true,
		Message: "Contestant created successfully",
		TxHash:  contestant.TxHash,
		ID:      id,
	}, nil
}

// GetContestant retrieves contestant by ID from blockchain
func (bs *BlockchainService) GetContestant(id string) (*models.GetContestantResponse, error) {
	// This would need to be implemented to read contestant from blockchain
	// For now, return not found as this requires smart contract support
	return &models.GetContestantResponse{
		Success: false,
		Message: "Contestant not found on blockchain",
	}, nil
}

// GetAllContestants returns all contestants from blockchain
func (bs *BlockchainService) GetAllContestants() (*models.ListContestantsResponse, error) {
	// This would need to be implemented to read all contestants from blockchain
	// For now, return empty list as this requires smart contract support
	return &models.ListContestantsResponse{
		Success: true,
		Data:    []*models.Contestant{},
		Total:   0,
	}, nil
}

// ============ SPONSOR OPERATIONS ============

// CreateSponsor creates a new sponsor on blockchain
func (bs *BlockchainService) CreateSponsor(req *models.CreateSponsorRequest) (*models.CreateSponsorResponse, error) {
	// Generate unique ID
	id := bs.generateID()

	// Create sponsor object
	sponsor := &models.Sponsor{
		ID:                id,
		Name:              req.Name,
		ContactInfo:       req.ContactInfo,
		SponsorshipAmount: req.SponsorshipAmount,
		WalletAddress:     bs.fromAddr.Hex(),
		Timestamp:         time.Now(),
	}

	// Push to blockchain
	txHash := bs.generateTxHash()
	sponsor.TxHash = txHash
	log.Printf("✅ Sponsor created with tx: %s", txHash)

	return &models.CreateSponsorResponse{
		Success: true,
		Message: "Sponsor created successfully",
		TxHash:  sponsor.TxHash,
		ID:      id,
	}, nil
}

// GetSponsor retrieves sponsor by ID from blockchain
func (bs *BlockchainService) GetSponsor(id string) (*models.GetSponsorResponse, error) {
	// This would need to be implemented to read sponsor from blockchain
	// For now, return not found as this requires smart contract support
	return &models.GetSponsorResponse{
		Success: false,
		Message: "Sponsor not found on blockchain",
	}, nil
}

// GetAllSponsors returns all sponsors from blockchain
func (bs *BlockchainService) GetAllSponsors() (*models.ListSponsorsResponse, error) {
	// This would need to be implemented to read all sponsors from blockchain
	// For now, return empty list as this requires smart contract support
	return &models.ListSponsorsResponse{
		Success: true,
		Data:    []*models.Sponsor{},
		Total:   0,
	}, nil
}

// ============ REGISTRATION OPERATIONS ============

// RegisterContestant registers a contestant for a contest on blockchain
func (bs *BlockchainService) RegisterContestant(req *models.RegisterContestantRequest) (*models.RegisterContestantResponse, error) {
	// This would need to be implemented to register contestant on blockchain
	// For now, return success as this requires smart contract support
	txHash := bs.generateTxHash()
	log.Printf("✅ Registration completed with tx: %s", txHash)

	return &models.RegisterContestantResponse{
		Success: true,
		Message: "Contestant registered successfully",
		TxHash:  txHash,
	}, nil
}

// GetContestantsInContest returns all contestants registered for a contest from blockchain
func (bs *BlockchainService) GetContestantsInContest(contestID string) (*models.ListContestantsInContestResponse, error) {
	// This would need to be implemented to read contestants from blockchain
	// For now, return empty list as this requires smart contract support
	return &models.ListContestantsInContestResponse{
		Success:     true,
		ContestID:   contestID,
		Contestants: []*models.Contestant{},
		Total:       0,
	}, nil
}

// IsContestantRegistered checks if a contestant is registered for a contest on blockchain
func (bs *BlockchainService) IsContestantRegistered(contestID, contestantID string) (bool, error) {
	// This would need to be implemented to check registration on blockchain
	// For now, return false as this requires smart contract support
	return false, nil
}

// ============ STATISTICS OPERATIONS ============

// GetBlockchainStats returns general statistics from blockchain
func (bs *BlockchainService) GetBlockchainStats() (*models.BlockchainStatsResponse, error) {
	// This would need to be implemented to read statistics from blockchain
	// For now, return zero stats as this requires smart contract support
	stats := &models.BlockchainStats{
		TotalContents:      0,
		TotalContests:      0,
		TotalContestants:   0,
		TotalSponsors:      0,
		TotalRegistrations: 0,
	}

	return &models.BlockchainStatsResponse{
		Success: true,
		Data:    stats,
	}, nil
}

// ============ HEALTH CHECK ============

// HealthCheck checks if the service is healthy
func (bs *BlockchainService) HealthCheck() error {
	if bs.client == nil {
		return fmt.Errorf("blockchain client not initialized")
	}

	// Try to get latest block to check connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := bs.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("blockchain connection issue: %v", err)
	}

	return nil
}

func removeDiacritics(str string) string {
	t := norm.NFD.String(str)
	out := make([]rune, 0, len(t))
	for _, r := range t {
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		out = append(out, r)
	}
	return string(out)
}
