# Blockchain Demo Tests

This directory contains comprehensive tests for the blockchain demo application.

## Test Structure

### 1. Real Blockchain Tests (`real_blockchain/`)
Tests that interact with a real blockchain network.

**Files:**
- `real_blockchain_test.go` - Tests for blockchain service functionality

**Test Cases:**
- `TestHealthCheck` - Tests blockchain connectivity
- `TestCreateAndGetContent` - Tests content creation and retrieval
- `TestCreateAndGetContest` - Tests contest creation and retrieval
- `TestSearchContests` - Tests search functionality
- `TestCreateContestant` - Tests contestant creation
- `TestCreateSponsor` - Tests sponsor creation
- `TestRegisterContestant` - Tests contestant registration
- `TestGetBlockchainStats` - Tests statistics retrieval
- `TestInvalidDateRange` - Tests error handling for invalid dates
- `TestInvalidDateFormat` - Tests error handling for invalid date formats

### 2. API Handler Tests (`internal/api/tests/`)
Tests for HTTP API endpoints.

**Files:**
- `handler_test.go` - Tests for API handlers

**Test Cases:**
- `TestHealthCheckHandler` - Tests health check endpoint
- `TestCreateContestHandler` - Tests contest creation endpoint
- `TestCreateContestHandlerInvalidRequest` - Tests error handling for invalid requests
- `TestSearchContestsHandler` - Tests search endpoint
- `TestSearchContestsHandlerMissingKeyword` - Tests error handling for missing parameters

### 3. Integration Tests (`integration_test.go`)
End-to-end tests that test the complete workflow.

**Test Cases:**
- `TestFullContestWorkflow` - Tests complete contest lifecycle
- `TestErrorHandling` - Tests error scenarios
- `TestSearchFunctionality` - Tests search with various keywords
- `TestConcurrentOperations` - Tests concurrent operations

## Running Tests

### Prerequisites

1. **Environment Setup:**
   ```bash
   # Copy environment file
   cp .env.example .env.test
   
   # Configure blockchain settings in .env.test
   NETWORK_URL=http://your-blockchain-node:8545
   CHAIN_ID=80001
   CONTRACT_ADDRESS=0x...
   PRIVATE_KEY=your_private_key_here
   ```

2. **Install Dependencies:**
   ```bash
   go mod tidy
   go get github.com/stretchr/testify/assert
   ```

### Running All Tests

```bash
# Run all tests
go test ./tests/... -v

# Run with coverage
go test ./tests/... -v -cover

# Run with race detection
go test ./tests/... -v -race
```

### Running Specific Test Categories

```bash
# Run only real blockchain tests
go test ./tests/real_blockchain/... -v

# Run only API handler tests
go test ./internal/api/tests/... -v

# Run only integration tests
go test ./tests/integration_test.go -v
```

### Running Individual Tests

```bash
# Run specific test function
go test ./tests/real_blockchain/... -v -run TestHealthCheck

# Run tests matching pattern
go test ./tests/real_blockchain/... -v -run "TestCreate.*"
```

## Test Configuration

### Environment Variables

The tests use the following environment variables (loaded from `.env.test`):

- `NETWORK_URL` - Blockchain network URL
- `CHAIN_ID` - Blockchain chain ID
- `CONTRACT_ADDRESS` - Deployed smart contract address
- `PRIVATE_KEY` - Private key for transactions

### Test Data

Tests create real transactions on the blockchain, so:
- Use a test network (e.g., Mumbai testnet)
- Ensure sufficient test tokens for gas fees
- Tests may take longer due to blockchain confirmations

## Test Coverage

### Current Coverage
- ✅ Blockchain service initialization
- ✅ Contest creation and validation
- ✅ Contestant creation
- ✅ Sponsor creation
- ✅ Registration functionality
- ✅ Search functionality
- ✅ Error handling
- ✅ API endpoints
- ✅ Health checks

### Areas for Improvement
- 🔄 Content retrieval (currently returns "not found")
- 🔄 Contest retrieval (currently returns "not found")
- 🔄 Statistics calculation (currently returns zeros)
- 🔄 Smart contract integration tests
- 🔄 Performance tests
- 🔄 Load tests

## Troubleshooting

### Common Issues

1. **Connection Errors:**
   ```
   failed to connect to blockchain
   ```
   - Check `NETWORK_URL` in `.env.test`
   - Ensure blockchain node is running
   - Check network connectivity

2. **Private Key Errors:**
   ```
   invalid private key
   ```
   - Verify `PRIVATE_KEY` format in `.env.test`
   - Ensure private key has sufficient funds

3. **Contract Errors:**
   ```
   Failed to load contract ABI
   ```
   - Ensure contract is deployed
   - Check `CONTRACT_ADDRESS` in `.env.test`
   - Verify ABI file exists at expected path

4. **Test Failures:**
   ```
   content not found on blockchain
   ```
   - This is expected for unimplemented functions
   - Tests are designed to handle these cases

### Debug Mode

Run tests with verbose output:
```bash
go test ./tests/... -v -debug
```

## Contributing

When adding new tests:

1. Follow the existing naming convention
2. Add proper error handling
3. Include both positive and negative test cases
4. Document any new test requirements
5. Update this README if adding new test categories

