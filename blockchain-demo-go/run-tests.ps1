# Script run-tests.ps1 để chạy test với blockchain thật
# Cấu hình kết nối đến mạng Hii Network

Write-Host "=== Chuẩn bị môi trường cho test với Hii Network ==="

# Đặt biến môi trường test - sử dụng Hii Network
$env:TEST_NETWORK_URL = "http://103.69.98.80:8545"
$env:TEST_CHAIN_ID = "22988"
$env:TEST_PRIVATE_KEY = "e6b362527d55d4e0a8b8330eb32b0e8eafed357a5dc73e3f8024a4f5d424edd7"
$env:TEST_CONTRACT_ADDRESS = "0xE3384d3f6794b8fa3Bdcb01De09245DF3893Ce45"

Write-Host "Biến môi trường đã được thiết lập:"
Write-Host "TEST_NETWORK_URL: $env:TEST_NETWORK_URL"
Write-Host "TEST_CHAIN_ID: $env:TEST_CHAIN_ID"
Write-Host "TEST_CONTRACT_ADDRESS: $env:TEST_CONTRACT_ADDRESS"
Write-Host "TEST_PRIVATE_KEY: [HIDDEN]"

# Tạo file .env.test
$envContent = @"
TEST_NETWORK_URL=http://103.69.98.80:8545
TEST_CHAIN_ID=22988
TEST_PRIVATE_KEY=e6b362527d55d4e0a8b8330eb32b0e8eafed357a5dc73e3f8024a4f5d424edd7
TEST_CONTRACT_ADDRESS=0xE3384d3f6794b8fa3Bdcb01De09245DF3893Ce45
"@

# Ghi file .env.test
$envContent | Out-File -FilePath ".env.test" -Encoding utf8

Write-Host "=== Chạy service test ==="
go test -v ./internal/service/...

Write-Host "=== Chạy API test ==="
go test -v ./internal/api/tests/...

# Tạo test report
$reportDir = "test-reports"
$reportFile = "$reportDir/test-report-" + (Get-Date -Format "yyyy-MM-dd_HH-mm-ss") + ".html"

# Tạo thư mục report nếu chưa tồn tại
if (!(Test-Path $reportDir)) {
    New-Item -ItemType Directory -Path $reportDir
}

# Chạy test với output là HTML report
Write-Host "=== Tạo test report ==="
go test -v ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o $reportFile

Write-Host "Test report được lưu tại: $reportFile"

# Xóa biến môi trường
Remove-Item Env:\TEST_NETWORK_URL
Remove-Item Env:\TEST_CHAIN_ID
Remove-Item Env:\TEST_PRIVATE_KEY
Remove-Item Env:\TEST_CONTRACT_ADDRESS

Write-Host "=== Hoàn thành! ===" 