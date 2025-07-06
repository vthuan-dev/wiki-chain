package main

import (
	"blockchain-demo/internal/api"
	"blockchain-demo/internal/config"
	"blockchain-demo/internal/service"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize blockchain service
	blockchainService, err := service.NewBlockchainService(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize blockchain service: %v", err)
	}

	// Initialize API handlers
	apiHandler := api.NewHandler(blockchainService)

	// Setup routes
	router := mux.NewRouter()

	// Add CORS middleware FIRST - with detailed logging
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			log.Printf("🌐 Request: %s %s from %s, Origin: %s", r.Method, r.URL.Path, r.RemoteAddr, origin)

			// Set CORS headers for all responses
			w.Header().Set("Access-Control-Allow-Origin", "*") // Cho phép tất cả origins
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Max-Age", "3600")

			// Handle preflight OPTIONS request
			if r.Method == "OPTIONS" {
				log.Printf("✓ Handling OPTIONS preflight request")
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	// API routes
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	// Test endpoint
	apiRouter.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("✅ Test endpoint hit!")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success":true,"message":"Test endpoint working!"}`))
	}).Methods("GET", "OPTIONS")

	// Add POST test endpoint specifically for testing contest creation
	apiRouter.HandleFunc("/test-post", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("🧪 Test POST endpoint called")

		// Read the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("❌ Error reading request body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Printf("📥 Received POST data: %s", string(body))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success":true,"message":"Test POST endpoint working!","txHash":"0xtest123","id":"test-id-123"}`))
	}).Methods("POST", "OPTIONS")

	// Content endpoints
	apiRouter.HandleFunc("/content", apiHandler.CreateContent).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/content/{id}", apiHandler.GetContent).Methods("GET", "OPTIONS")
	apiRouter.HandleFunc("/contents", apiHandler.ListContents).Methods("GET", "OPTIONS")

	// Fix Contest endpoints by using explicit subrouter for method separation
	contestsRouter := apiRouter.PathPrefix("/contests").Subrouter()
	contestsRouter.HandleFunc("", apiHandler.ListContests).Methods("GET", "OPTIONS")    // GET /contests
	contestsRouter.HandleFunc("", apiHandler.CreateContest).Methods("POST", "OPTIONS")  // POST /contests
	contestsRouter.HandleFunc("/{id}", apiHandler.GetContest).Methods("GET", "OPTIONS") // GET /contests/{id}
	contestsRouter.HandleFunc("/{contestId}/register", apiHandler.RegisterContestant).Methods("POST", "OPTIONS")
	contestsRouter.HandleFunc("/{contestId}/contestants", apiHandler.GetContestantsInContest).Methods("GET", "OPTIONS")

	// Contestant endpoints
	apiRouter.HandleFunc("/contestants", apiHandler.CreateContestant).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/contestants/{id}", apiHandler.GetContestant).Methods("GET", "OPTIONS")
	apiRouter.HandleFunc("/contestants", apiHandler.ListContestants).Methods("GET", "OPTIONS")

	// Sponsor endpoints
	apiRouter.HandleFunc("/sponsors", apiHandler.CreateSponsor).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/sponsors/{id}", apiHandler.GetSponsor).Methods("GET", "OPTIONS")
	apiRouter.HandleFunc("/sponsors", apiHandler.ListSponsors).Methods("GET", "OPTIONS")

	// Statistics endpoint
	apiRouter.HandleFunc("/stats", apiHandler.GetStats).Methods("GET", "OPTIONS")

	// Health check
	apiRouter.HandleFunc("/health", apiHandler.HealthCheck).Methods("GET", "OPTIONS")

	log.Printf("🚀 Server starting on %s:%s", cfg.Host, cfg.Port)
	log.Printf("📡 Connected to blockchain: %s", cfg.NetworkURL)
	log.Printf("📋 API Documentation available at: http://%s:%s/api/v1", cfg.Host, cfg.Port)

	if err := http.ListenAndServe(cfg.Host+":"+cfg.Port, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
