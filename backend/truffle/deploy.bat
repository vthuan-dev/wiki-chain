@echo off
echo 🚀 Bắt đầu deploy smart contract lên Hii Network...

REM Kiểm tra file .env
if not exist .env (
    echo ❌ File .env không tồn tại! Tạo file .env từ .env.example
    exit /b 1
)

echo 📋 Kiểm tra cấu hình...
echo Network: Hii Network ^(Chain ID: 22988^)
echo RPC URL: http://103.69.98.80:8545
echo.

REM Kiểm tra Node.js dependencies
echo 📦 Kiểm tra dependencies...
if not exist node_modules (
    echo 📥 Installing dependencies...
    call npm install
)

REM Compile contracts
echo 🔨 Compiling contracts...
call npx truffle compile

if %errorlevel% neq 0 (
    echo ❌ Compile failed!
    pause
    exit /b 1
)

echo ✅ Compile thành công!
echo.

REM Deploy to Hii Network
echo 🚀 Deploying to Hii Network...
call npx truffle migrate --network hii --reset

if %errorlevel% equ 0 (
    echo.
    echo ✅ Deploy thành công!
    echo.
    
    REM Update Go backend configuration
    echo 🔄 Updating Go backend configuration...
    call node scripts\update-go-config.js
    
    echo.
    echo 🎉 Deployment completed successfully!
    echo.
    echo 📋 Summary:
    echo    ✅ Contract deployed to Hii Network
    echo    ✅ Go backend configuration updated
    echo    🔍 Check contract on: https://explorer.testnet.hii.network
    echo.
    echo 📝 Next steps:
    echo    1. cd ..\..\blockchain-demo-go
    echo    2. go run cmd\server\main.go
    echo    3. Test API at: http://localhost:8080
    
) else (
    echo.
    echo ❌ Deploy thất bại!
    echo 💡 Kiểm tra:
    echo    - Số dư ví đủ gas
    echo    - Network connection
    echo    - Private key đúng
)

echo.
pause
