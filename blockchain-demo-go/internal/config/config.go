// internal/config/config.go
package config

import (
	"os"
	"path/filepath"

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
	ContractJSON    string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {

	_ = godotenv.Load()


	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}


	defaultContractPath := filepath.Join(pwd, "..", "backend", "truffle", "build", "contracts", "ContentStorage.json")

	config := &Config{
		Host: getEnv("HOST", "localhost"),
		Port: getEnv("PORT", "8081"),

		NetworkURL:      getEnv("NETWORK_URL", "https://rpc.ankr.com/polygon_mumbai"),
		ChainID:         getEnv("CHAIN_ID", "80001"),
		ContractAddress: getEnv("CONTRACT_ADDRESS", ""),
		PrivateKey:      getEnv("PRIVATE_KEY", ""),
		ContractJSON:    getEnv("CONTRACT_JSON", defaultContractPath),
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
