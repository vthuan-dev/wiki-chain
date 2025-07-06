# API Documentation - Event Management Blockchain Demo

## 📚 Tổng quan API

API này cung cấp các endpoint để quản lý sự kiện, thí sinh, nhà tài trợ và đăng ký tham gia trên blockchain.

**Base URL**: `http://localhost:8080/api/v1`

---

## 🎯 1. Content APIs (Nội dung chung)

### 1.1 Tạo nội dung
```http
POST /api/v1/content
```

**Request Body:**
```json
{
  "title": "Hướng dẫn tham gia cuộc thi Board Game",
  "content": "Các bước cần thiết để tham gia cuộc thi thiết kế board game Việt Nam 2025...",
  "creator": "admin"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Content created successfully",
  "tx_hash": "0x1234567890abcdef...",
  "id": "content_12345"
}
```

### 1.2 Lấy nội dung
```http
GET /api/v1/content/{id}
```

### 1.3 Danh sách nội dung
```http
GET /api/v1/contents
```

---

## 🏆 2. Contest APIs (Cuộc thi)

### 2.1 Tạo cuộc thi mới
```http
POST /api/v1/contests
```

**Request Body:**
```json
{
  "name": "Road To ESSEN 2025",
  "description": "Cuộc thi thiết kế board game Việt Nam lớn nhất năm 2025",
  "start_date": "2025-06-24T00:00:00Z",
  "end_date": "2025-07-20T23:59:59Z",
  "image_url": "https://example.com/contest-image.jpg"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Contest created successfully",
  "tx_hash": "0xabcdef1234567890...",
  "id": "contest_67890"
}
```

### 2.2 Lấy thông tin cuộc thi
```http
GET /api/v1/contests/{id}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": "contest_67890",
    "name": "Road To ESSEN 2025",
    "description": "Cuộc thi thiết kế board game...",
    "start_date": "2025-06-24T00:00:00Z",
    "end_date": "2025-07-20T23:59:59Z",
    "organizer": "0x742d35cc6641c7b2b85ce462af7c9bb7a5db8b7a",
    "active": true,
    "image_url": "https://example.com/contest-image.jpg",
    "timestamp": "2025-07-05T10:30:00Z",
    "tx_hash": "0xabcdef1234567890..."
  }
}
```

### 2.3 Danh sách tất cả cuộc thi
```http
GET /api/v1/contests
```

---

## 👤 3. Contestant APIs (Thí sinh)

### 3.1 Đăng ký thí sinh
```http
POST /api/v1/contestants
```

**Request Body:**
```json
{
  "name": "Nguyễn Văn A",
  "details": "Sinh viên năm 4 chuyên ngành Thiết kế Game, có kinh nghiệm 2 năm làm board game indie",
  "creator": "0x742d35cc6641c7b2b85ce462af7c9bb7a5db8b7a"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Contestant created successfully",
  "tx_hash": "0xfedcba0987654321...",
  "id": "contestant_11111"
}
```

### 3.2 Lấy thông tin thí sinh
```http
GET /api/v1/contestants/{id}
```

### 3.3 Danh sách thí sinh
```http
GET /api/v1/contestants
```

---

## 💰 4. Sponsor APIs (Nhà tài trợ)

### 4.1 Đăng ký nhà tài trợ
```http
POST /api/v1/sponsors
```

**Request Body:**
```json
{
  "name": "VNG Corporation",
  "contact_info": "sponsor@vng.com.vn | 028-1234-5678",
  "sponsorship_amount": 50000000
}
```

**Response:**
```json
{
  "success": true,
  "message": "Sponsor created successfully",
  "tx_hash": "0x9876543210fedcba...",
  "id": "sponsor_22222"
}
```

### 4.2 Lấy thông tin nhà tài trợ
```http
GET /api/v1/sponsors/{id}
```

### 4.3 Danh sách nhà tài trợ
```http
GET /api/v1/sponsors
```

---

## 📝 5. Registration APIs (Đăng ký tham gia)

### 5.1 Đăng ký thí sinh vào cuộc thi
```http
POST /api/v1/contests/{contestId}/register
```

**Request Body:**
```json
{
  "contestant_id": "contestant_11111"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Contestant registered successfully",
  "tx_hash": "0x5555666677778888..."
}
```

### 5.2 Lấy danh sách thí sinh trong cuộc thi
```http
GET /api/v1/contests/{contestId}/contestants
```

**Response:**
```json
{
  "success": true,
  "contest_id": "contest_67890",
  "contestants": [
    {
      "id": "contestant_11111",
      "name": "Nguyễn Văn A",
      "details": "Sinh viên năm 4...",
      "creator": "0x742d35cc6641c7b2b85ce462af7c9bb7a5db8b7a",
      "timestamp": "2025-07-05T10:45:00Z",
      "verified": true,
      "tx_hash": "0xfedcba0987654321..."
    }
  ],
  "total": 1
}
```

---

## 📊 6. Statistics API (Thống kê)

### 6.1 Lấy thống kê tổng quan
```http
GET /api/v1/stats
```

**Response:**
```json
{
  "success": true,
  "data": {
    "total_contents": 5,
    "total_contests": 3,
    "total_contestants": 15,
    "total_sponsors": 8,
    "total_registrations": 45
  }
}
```

---

## ❤️ 7. Health Check

### 7.1 Kiểm tra trạng thái hệ thống
```http
GET /api/v1/health
```

**Response:**
```json
{
  "status": "healthy",
  "blockchain": "connected",
  "message": "Service is running properly"
}
```

---

## 🚀 Ví dụ workflow hoàn chỉnh

### Bước 1: Tạo cuộc thi
```bash
curl -X POST http://localhost:8080/api/v1/contests \
  -H "Content-Type: application/json" \
  -d '{
    "name": "THE MC FACE 2025",
    "description": "Cuộc thi tìm kiếm MC tài năng cho sinh viên",
    "start_date": "2025-06-29T00:00:00Z",
    "end_date": "2025-08-15T23:59:59Z",
    "image_url": "https://example.com/mc-face.jpg"
  }'
```

### Bước 2: Đăng ký thí sinh
```bash
curl -X POST http://localhost:8080/api/v1/contestants \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Trần Thị B",
    "details": "Sinh viên trường Đại học Văn Lang, có kinh nghiệm MC sự kiện học đường",
    "creator": "0x123456789abcdef..."
  }'
```

### Bước 3: Đăng ký thí sinh vào cuộc thi
```bash
curl -X POST http://localhost:8080/api/v1/contests/{contest_id}/register \
  -H "Content-Type: application/json" \
  -d '{
    "contestant_id": "{contestant_id}"
  }'
```

### Bước 4: Kiểm tra danh sách thí sinh đã đăng ký
```bash
curl http://localhost:8080/api/v1/contests/{contest_id}/contestants
```

---

## 🔧 Cài đặt và chạy

1. **Clone dự án**
2. **Cài đặt dependencies**: `go mod tidy`
3. **Cấu hình .env** (sao chép từ .env.example)
4. **Chạy server**: `go run cmd/server/main.go`
5. **API sẽ chạy tại**: http://localhost:8080

---

## 📝 Lưu ý

- Tất cả timestamps sử dụng format RFC3339: `2006-01-02T15:04:05Z`
- API hỗ trợ CORS cho frontend integration
- Tất cả response đều có format JSON
- Error responses có cấu trúc thống nhất với `success: false`
