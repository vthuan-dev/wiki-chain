package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	// Server configuration
	Host string
	Port string

	// Blockchain configuration
	NetworkURL      string
	ChainID         string
	ContractAddress string
	PrivateKey      string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists (optional)
	_ = godotenv.Load()

	config := &Config{
		Host: getEnv("HOST", "localhost"),
		Port: getEnv("PORT", "8081"), // Đổi từ 8080 sang 8081

		NetworkURL:      getEnv("NETWORK_URL", "https://rpc.ankr.com/polygon_mumbai"),
		ChainID:         getEnv("CHAIN_ID", "80001"),
		ContractAddress: getEnv("CONTRACT_ADDRESS", ""),
		PrivateKey:      getEnv("PRIVATE_KEY", ""),
	}

	return config, nil
}

// getEnv gets environment variable with default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
