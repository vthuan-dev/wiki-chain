@echo off

echo 🚀 Starting Blockchain Demo Setup...

REM Check if Go is installed
go version >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ Go is not installed. Please install Go first.
    exit /b 1
)

echo ✅ Go is installed
go version

REM Navigate to project directory
cd /d "%~dp0"

echo 📦 Installing dependencies...
go mod tidy

echo 🔧 Creating .env file from template...
if not exist .env (
    copy .env.example .env
    echo 📝 Please edit .env file with your blockchain configuration
) else (
    echo ⚠️  .env file already exists
)

echo 🏗️  Building application...
if not exist bin mkdir bin
go build -o bin\server.exe .\cmd\server

echo ✅ Setup completed!
echo.
echo To run the application:
echo   bin\server.exe
echo or
echo   go run cmd\server\main.go
echo.
echo API will be available at: http://localhost:8080

pause
