# Hướng dẫn Deploy Smart Contract và Chạy Backend



## 🔧 Bước 1: Deploy Smart Contract

### Windows:
```bash
cd d:\intern\test-ui\backend\truffle
.\deploy.bat
```

### Linux/Mac:
```bash
cd /d/intern/test-ui/backend/truffle
chmod +x deploy.sh
./deploy.sh
```

## ✅ Kết quả mong đợi:
```
🚀 Bắt đầu deploy smart contract lên Hii Network...
📋 Kiểm tra cấu hình...
🔨 Compiling contracts...
✅ Compile thành công!
🚀 Deploying to Hii Network...

Deploying 'ContentStorage'
---------------------------
> transaction hash:    0xabc123...
> contract address:    0x1234567890abcdef...
> block number:        123456
> gas used:            2,345,678

✅ Deploy thành công!
🔄 Updating Go backend configuration...
✅ Đã cập nhật Go backend configuration
📄 Contract Address: 0x1234567890abcdef...

🎉 Deployment completed successfully!
```

## 🔧 Bước 2: Chạy Go Backend

```bash
cd d:\intern\test-ui\blockchain-demo-go
go run cmd\server\main.go
```

### Kết quả mong đợi:
```
✅ Loaded wallet address: 0xB10bd1778a4373E0249B9121DC64ab5285Fb3e1F
✅ Connected to blockchain network: http://103.69.98.80:8545
🚀 Server starting on localhost:8080
📡 Connected to blockchain: http://103.69.98.80:8545
📋 API Documentation available at: http://localhost:8080/api/v1
```

## 🧪 Bước 3: Test API

### PowerShell:
```powershell
cd d:\intern\test-ui\blockchain-demo-go
.\test-api.ps1
```

### Curl:
```bash
# Health check
curl http://localhost:8080/api/v1/health

# Tạo contest
curl -X POST http://localhost:8080/api/v1/contests \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Road To ESSEN 2025",
    "description": "Cuộc thi thiết kế board game Việt Nam",
    "start_date": "2025-06-24T00:00:00Z",
    "end_date": "2025-07-20T23:59:59Z"
  }'
```

## 🔍 Bước 4: Verify Contract trên Explorer

1. Mở: https://explorer.testnet.hii.network
2. Tìm kiếm contract address từ output deploy
3. Kiểm tra transactions và contract details

## 📱 Bước 5: Tích hợp Frontend

Contract address sẽ được tự động cập nhật vào:
- `blockchain-demo-go/.env`
- Frontend có thể call API tại `http://localhost:8080/api/v1`

## 🔧 Troubleshooting

### Lỗi: "insufficient funds"
- Kiểm tra số dư ví trên Hii Network
- Request faucet nếu cần

### Lỗi: "connection refused"
- Kiểm tra RPC URL: http://103.69.98.80:8545
- Thử ping network

### Lỗi: "nonce too low"
- Reset nonce trong MetaMask
- Hoặc dùng `--reset` flag

## 📋 Các file quan trọng:

- `backend/truffle/.env` - Private key deployment
- `backend/truffle/truffle-config.js` - Network config  
- `blockchain-demo-go/.env` - Go backend config
- `backend/truffle/build/contracts/ContentStorage.json` - Contract ABI

## 🔐 Bảo mật:

1. **KHÔNG commit .env files**
2. **Sử dụng ví mới cho production**
3. **Backup private key an toàn**
4. **Test trên testnet trước**
