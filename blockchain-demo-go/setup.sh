#!/bin/bash

echo "🚀 Starting Blockchain Demo Setup..."

# Check if Go is installed
if ! command -v go &> /dev/null
then
    echo "❌ Go is not installed. Please install Go first."
    exit 1
fi

echo "✅ Go is installed: $(go version)"

# Navigate to project directory
cd "$(dirname "$0")"

echo "📦 Installing dependencies..."
go mod tidy

echo "🔧 Creating .env file from template..."
if [ ! -f .env ]; then
    cp .env.example .env
    echo "📝 Please edit .env file with your blockchain configuration"
else
    echo "⚠️  .env file already exists"
fi

echo "🏗️  Building application..."
go build -o bin/server ./cmd/server

echo "✅ Setup completed!"
echo ""
echo "To run the application:"
echo "  ./bin/server"
echo "or"
echo "  go run cmd/server/main.go"
echo ""
echo "API will be available at: http://localhost:8080"
