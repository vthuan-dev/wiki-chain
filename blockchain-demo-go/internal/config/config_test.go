// internal/config/config_test.go

package config

import (
    "os"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
    // Backup existing env vars
    oldNetworkURL := os.Getenv("NETWORK_URL")
    oldChainID := os.Getenv("CHAIN_ID")
    oldPrivateKey := os.Getenv("PRIVATE_KEY")
    oldContractAddr := os.Getenv("CONTRACT_ADDRESS")

    // Restore env vars after test
    defer func() {
        os.Setenv("NETWORK_URL", oldNetworkURL)
        os.Setenv("CHAIN_ID", oldChainID)
        os.Setenv("PRIVATE_KEY", oldPrivateKey)
        os.Setenv("CONTRACT_ADDRESS", oldContractAddr)
    }()

    // Set test env vars
    os.Setenv("NETWORK_URL", "http://localhost:8545")
    os.Setenv("CHAIN_ID", "1337")
    os.Setenv("PRIVATE_KEY", "test_key")
    os.Setenv("CONTRACT_ADDRESS", "0x123")

    cfg, err := Load()
    assert.NoError(t, err)
    assert.NotNil(t, cfg)
    assert.Equal(t, "http://localhost:8545", cfg.NetworkURL)
    assert.Equal(t, "1337", cfg.ChainID)
    assert.Equal(t, "test_key", cfg.PrivateKey)
    assert.Equal(t, "0x123", cfg.ContractAddress)
}

func TestGetEnv(t *testing.T) {
    // Test with existing env var
    os.Setenv("TEST_VAR", "test_value")
    value := getEnv("TEST_VAR", "default")
    assert.Equal(t, "test_value", value)

    // Test with non-existing env var
    value = getEnv("NON_EXISTING_VAR", "default")
    assert.Equal(t, "default", value)
}