#!/bin/bash

echo "🚀 Bắt đầu deploy smart contract lên Hii Network..."

# Kiểm tra môi trường
if [ ! -f .env ]; then
    echo "❌ File .env không tồn tại! Tạo file .env từ .env.example"
    exit 1
fi

echo "📋 Kiểm tra cấu hình..."
echo "Network: Hii Network (Chain ID: 22988)"
echo "RPC URL: http://103.69.98.80:8545"

# Compile contracts
echo "🔨 Compiling contracts..."
npx truffle compile

if [ $? -ne 0 ]; then
    echo "❌ Compile failed!"
    exit 1
fi

# Deploy to Hii Network
echo "🚀 Deploying to Hii Network..."
npx truffle migrate --network hii --reset

if [ $? -eq 0 ]; then
    echo "✅ Deploy thành công!"
    echo "📄 Kiểm tra contract address trong output ở trên"
    echo "🔍 Verify contract trên: https://explorer.testnet.hii.network"
else
    echo "❌ Deploy thất bại!"
    exit 1
fi
