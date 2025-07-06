package api

import (
	"blockchain-demo/internal/models"
	"blockchain-demo/internal/service"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler handles HTTP requests
type Handler struct {
	blockchainService *service.BlockchainService
}

// NewHandler creates a new API handler
func NewHandler(blockchainService *service.BlockchainService) *Handler {
	return &Handler{
		blockchainService: blockchainService,
	}
}

// CreateContent handles POST /api/v1/content
func (h *Handler) CreateContent(w http.ResponseWriter, r *http.Request) {
	var req models.CreateContentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	// Validate required fields
	if req.Title == "" {
		h.respondWithError(w, http.StatusBadRequest, "Title is required", "")
		return
	}

	if req.Content == "" {
		h.respondWithError(w, http.StatusBadRequest, "Content is required", "")
		return
	}

	// Set default creator if not provided
	if req.Creator == "" {
		req.Creator = "anonymous"
	}

	log.Printf("📝 Creating content: %s by %s", req.Title, req.Creator)

	// Create content via blockchain service
	response, err := h.blockchainService.StoreContent(&req)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to create content", err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusCreated, response)
}

// GetContent handles GET /api/v1/content/{id}
func (h *Handler) GetContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		h.respondWithError(w, http.StatusBadRequest, "Content ID is required", "")
		return
	}

	log.Printf("📖 Getting content: %s", id)

	// Get content via blockchain service
	response, err := h.blockchainService.GetContent(id)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to get content", err.Error())
		return
	}

	if !response.Success {
		h.respondWithJSON(w, http.StatusNotFound, response)
		return
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

// ListContents handles GET /api/v1/contents
func (h *Handler) ListContents(w http.ResponseWriter, r *http.Request) {
	log.Printf("📋 Listing all contents")

	// Get all contents via blockchain service
	response, err := h.blockchainService.GetAllContents()
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to list contents", err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

// HealthCheck handles GET /api/v1/health
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Check blockchain connection
	if err := h.blockchainService.HealthCheck(); err != nil {
		response := map[string]interface{}{
			"status":     "unhealthy",
			"blockchain": "disconnected",
			"error":      err.Error(),
		}
		h.respondWithJSON(w, http.StatusServiceUnavailable, response)
		return
	}

	response := map[string]interface{}{
		"status":     "healthy",
		"blockchain": "connected",
		"message":    "Service is running properly",
	}
	h.respondWithJSON(w, http.StatusOK, response)
}

// ============ CONTEST HANDLERS ============

// CreateContest handles POST /api/v1/contests
func (h *Handler) CreateContest(w http.ResponseWriter, r *http.Request) {
	log.Printf("📝 CreateContest called - Method: %s, URL: %s", r.Method, r.URL.Path)

	var req models.CreateContestRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("❌ Error decoding request body: %v", err)
		h.respondWithError(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	log.Printf("📋 Request data: %+v", req)

	// Validate required fields
	if req.Name == "" {
		h.respondWithError(w, http.StatusBadRequest, "Contest name is required", "")
		return
	}

	if req.Description == "" {
		h.respondWithError(w, http.StatusBadRequest, "Contest description is required", "")
		return
	}

	if req.StartDate == "" || req.EndDate == "" {
		h.respondWithError(w, http.StatusBadRequest, "Start date and end date are required", "")
		return
	}

	log.Printf("🏆 Creating contest: %s", req.Name)

	response, err := h.blockchainService.CreateContest(&req)
	if err != nil {
		log.Printf("❌ Error creating contest: %v", err)
		h.respondWithError(w, http.StatusInternalServerError, "Failed to create contest", err.Error())
		return
	}

	log.Printf("✅ Contest created successfully: %+v", response)
	h.respondWithJSON(w, http.StatusCreated, response)
}

// GetContest handles GET /api/v1/contests/{id}
func (h *Handler) GetContest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		h.respondWithError(w, http.StatusBadRequest, "Contest ID is required", "")
		return
	}

	log.Printf("🏆 Getting contest: %s", id)

	response, err := h.blockchainService.GetContest(id)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to get contest", err.Error())
		return
	}

	if !response.Success {
		h.respondWithJSON(w, http.StatusNotFound, response)
		return
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

// ListContests handles GET /api/v1/contests
func (h *Handler) ListContests(w http.ResponseWriter, r *http.Request) {
	log.Printf("🏆 Listing all contests")

	response, err := h.blockchainService.GetAllContests()
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to list contests", err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

// ============ CONTESTANT HANDLERS ============

// CreateContestant handles POST /api/v1/contestants
func (h *Handler) CreateContestant(w http.ResponseWriter, r *http.Request) {
	var req models.CreateContestantRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	// Validate required fields
	if req.Name == "" {
		h.respondWithError(w, http.StatusBadRequest, "Contestant name is required", "")
		return
	}

	if req.Details == "" {
		h.respondWithError(w, http.StatusBadRequest, "Contestant details are required", "")
		return
	}

	// Set default creator if not provided
	if req.Creator == "" {
		req.Creator = "anonymous"
	}

	log.Printf("👤 Creating contestant: %s", req.Name)

	response, err := h.blockchainService.CreateContestant(&req)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to create contestant", err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusCreated, response)
}

// GetContestant handles GET /api/v1/contestants/{id}
func (h *Handler) GetContestant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		h.respondWithError(w, http.StatusBadRequest, "Contestant ID is required", "")
		return
	}

	log.Printf("👤 Getting contestant: %s", id)

	response, err := h.blockchainService.GetContestant(id)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to get contestant", err.Error())
		return
	}

	if !response.Success {
		h.respondWithJSON(w, http.StatusNotFound, response)
		return
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

// ListContestants handles GET /api/v1/contestants
func (h *Handler) ListContestants(w http.ResponseWriter, r *http.Request) {
	log.Printf("👤 Listing all contestants")

	response, err := h.blockchainService.GetAllContestants()
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to list contestants", err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

// ============ SPONSOR HANDLERS ============

// CreateSponsor handles POST /api/v1/sponsors
func (h *Handler) CreateSponsor(w http.ResponseWriter, r *http.Request) {
	var req models.CreateSponsorRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	// Validate required fields
	if req.Name == "" {
		h.respondWithError(w, http.StatusBadRequest, "Sponsor name is required", "")
		return
	}

	if req.ContactInfo == "" {
		h.respondWithError(w, http.StatusBadRequest, "Contact info is required", "")
		return
	}

	log.Printf("💰 Creating sponsor: %s", req.Name)

	response, err := h.blockchainService.CreateSponsor(&req)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to create sponsor", err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusCreated, response)
}

// GetSponsor handles GET /api/v1/sponsors/{id}
func (h *Handler) GetSponsor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		h.respondWithError(w, http.StatusBadRequest, "Sponsor ID is required", "")
		return
	}

	log.Printf("💰 Getting sponsor: %s", id)

	response, err := h.blockchainService.GetSponsor(id)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to get sponsor", err.Error())
		return
	}

	if !response.Success {
		h.respondWithJSON(w, http.StatusNotFound, response)
		return
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

// ListSponsors handles GET /api/v1/sponsors
func (h *Handler) ListSponsors(w http.ResponseWriter, r *http.Request) {
	log.Printf("💰 Listing all sponsors")

	response, err := h.blockchainService.GetAllSponsors()
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to list sponsors", err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

// ============ REGISTRATION HANDLERS ============

// RegisterContestant handles POST /api/v1/contests/{contestId}/register
func (h *Handler) RegisterContestant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contestID := vars["contestId"]

	if contestID == "" {
		h.respondWithError(w, http.StatusBadRequest, "Contest ID is required", "")
		return
	}

	var req models.RegisterContestantRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	// Override contest ID from URL
	req.ContestID = contestID

	if req.ContestantID == "" {
		h.respondWithError(w, http.StatusBadRequest, "Contestant ID is required", "")
		return
	}

	log.Printf("📝 Registering contestant %s for contest %s", req.ContestantID, req.ContestID)

	response, err := h.blockchainService.RegisterContestant(&req)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to register contestant", err.Error())
		return
	}

	if !response.Success {
		h.respondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	h.respondWithJSON(w, http.StatusCreated, response)
}

// GetContestantsInContest handles GET /api/v1/contests/{contestId}/contestants
func (h *Handler) GetContestantsInContest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contestID := vars["contestId"]

	if contestID == "" {
		h.respondWithError(w, http.StatusBadRequest, "Contest ID is required", "")
		return
	}

	log.Printf("👥 Getting contestants for contest: %s", contestID)

	response, err := h.blockchainService.GetContestantsInContest(contestID)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to get contestants", err.Error())
		return
	}

	if !response.Success {
		h.respondWithJSON(w, http.StatusNotFound, response)
		return
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

// ============ STATISTICS HANDLERS ============

// GetStats handles GET /api/v1/stats
func (h *Handler) GetStats(w http.ResponseWriter, r *http.Request) {
	log.Printf("📊 Getting blockchain statistics")

	response, err := h.blockchainService.GetBlockchainStats()
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to get statistics", err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

// SearchContestsHandler handles GET /api/v1/contests/search?keyword=abc
func (h *Handler) SearchContestsHandler(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	if keyword == "" {
		h.respondWithError(w, http.StatusBadRequest, "Missing keyword", "")
		return
	}
	results, err := h.blockchainService.SearchContests(keyword)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Search failed", err.Error())
		return
	}
	h.respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    results,
		"total":   len(results),
	})
}

// respondWithError sends an error response
func (h *Handler) respondWithError(w http.ResponseWriter, code int, message, details string) {
	response := models.ErrorResponse{
		Success: false,
		Error:   message,
		Message: details,
	}
	h.respondWithJSON(w, code, response)
}

// respondWithJSON sends a JSON response
func (h *Handler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	log.Printf("📤 Sending JSON response - Status: %d, Payload: %+v", code, payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("❌ Failed to encode JSON response: %v", err)
	} else {
		log.Printf("✅ JSON response sent successfully")
	}
}
