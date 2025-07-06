# Blockchain Demo Go - Event Management System

Ứng dụng demo blockchain đơn giản được viết bằng Golang, tương thích với hệ thống quản lý sự kiện như website votome.now của bạn.

## 🚀 Tính năng chính

- ✅ **Quản lý cuộc thi**: Tạo, xem và quản lý các cuộc thi/sự kiện
- ✅ **Quản lý thí sinh**: Đăng ký và quản lý thông tin thí sinh
- ✅ **Quản lý nhà tài trợ**: Đăng ký và theo dõi nhà tài trợ
- ✅ **Đăng ký tham gia**: Cho phép thí sinh đăng ký vào cuộc thi
- ✅ **Push/Get dữ liệu blockchain**: Lưu trữ và truy xuất dữ liệu từ blockchain
- ✅ **REST API**: API endpoints đầy đủ cho frontend integration
- ✅ **EVM Compatible**: Hỗ trợ các mạng blockchain tương thích EVM
- ✅ **Demo Mode**: Có thể chạy mà không cần kết nối blockchain thật

## 📁 Cấu trúc dự án

```
blockchain-demo-go/
├── cmd/
│   └── server/
│       └── main.go              # Entry point của ứng dụng
├── internal/
│   ├── api/
│   │   └── handler.go           # HTTP handlers
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── models/
│   │   └── content.go           # Data structures
│   └── service/
│       └── blockchain.go        # Blockchain logic
├── .env.example                 # Environment variables template
├── go.mod                       # Go modules
└── README.md                    # Documentation
```

## 🛠️ Cài đặt và chạy

### 1. Cài đặt dependencies

```bash
cd blockchain-demo-go
go mod tidy
```

### 2. Cấu hình environment

```bash
cp .env.example .env
```

Chỉnh sửa file `.env`:
```env
# Blockchain Network Settings
NETWORK_URL=https://rpc.ankr.com/polygon_mumbai
CHAIN_ID=80001

# Contract Settings (optional)
CONTRACT_ADDRESS=0x1234567890123456789012345678901234567890

# Private Key (optional - for real blockchain interaction)
PRIVATE_KEY=your_private_key_here

# Server Settings
PORT=8080
HOST=localhost
```

### 3. Chạy ứng dụng

```bash
go run cmd/server/main.go
```

Server sẽ chạy tại: `http://localhost:8080`

## 📡 API Endpoints

### 1. Health Check
```http
GET /api/v1/health
```

### 2. Tạo nội dung (Push to Blockchain)
```http
POST /api/v1/content
Content-Type: application/json

{
  "title": "Tiêu đề nội dung",
  "content": "Nội dung chi tiết...",
  "creator": "địa chỉ ví hoặc tên người tạo"
}
```

### 3. Lấy nội dung (Get from Blockchain)
```http
GET /api/v1/content/{id}
```

### 4. Liệt kê tất cả nội dung
```http
GET /api/v1/contents
```

## 🧪 Test API

### Sử dụng PowerShell script
```powershell
.\test-api.ps1
```

### Hoặc test thủ công với curl

#### Tạo cuộc thi mới
```bash
curl -X POST http://localhost:8080/api/v1/contests \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Road To ESSEN 2025",
    "description": "Cuộc thi thiết kế board game Việt Nam lớn nhất năm 2025",
    "start_date": "2025-06-24T00:00:00Z",
    "end_date": "2025-07-20T23:59:59Z",
    "image_url": "https://example.com/contest.jpg"
  }'
```

#### Đăng ký thí sinh
```bash
curl -X POST http://localhost:8080/api/v1/contestants \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Nguyễn Văn A", 
    "details": "Sinh viên chuyên ngành Game Design",
    "creator": "0x742d35cc6641c7b2b85ce462af7c9bb7a5db8b7a"
  }'
```

#### Đăng ký thí sinh vào cuộc thi
```bash
curl -X POST http://localhost:8080/api/v1/contests/{contest_id}/register \
  -H "Content-Type: application/json" \
  -d '{
    "contestant_id": "{contestant_id}"
  }'
```

## 🔗 Tích hợp với Frontend React

Để tích hợp với dự án React frontend hiện tại của bạn:

### 1. Tạo API service trong frontend
```typescript
// src/services/blockchainApi.ts
const API_BASE = 'http://localhost:8080/api/v1';

export const blockchainApi = {
  // Contest APIs
  createContest: async (data: CreateContestRequest) => {
    const response = await fetch(`${API_BASE}/contests`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data)
    });
    return response.json();
  },

  getContests: async () => {
    const response = await fetch(`${API_BASE}/contests`);
    return response.json();
  },

  // Contestant APIs
  createContestant: async (data: CreateContestantRequest) => {
    const response = await fetch(`${API_BASE}/contestants`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data)
    });
    return response.json();
  },

  // Registration APIs
  registerContestant: async (contestId: string, contestantId: string) => {
    const response = await fetch(`${API_BASE}/contests/${contestId}/register`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ contestant_id: contestantId })
    });
    return response.json();
  }
};
```

### 2. Sử dụng trong React components
```tsx
// src/components/CreateEventForm.tsx
import { blockchainApi } from '@/services/blockchainApi';

const CreateEventForm = () => {
  const handleSubmit = async (formData) => {
    try {
      const result = await blockchainApi.createContest({
        name: formData.title,
        description: formData.description,
        start_date: formData.startDate,
        end_date: formData.endDate,
        image_url: formData.imageUrl
      });
      
      if (result.success) {
        console.log('✅ Contest created on blockchain:', result.tx_hash);
        toast.success('Sự kiện đã được tạo và lưu trên blockchain!');
      }
    } catch (error) {
      console.error('❌ Error:', error);
      toast.error('Có lỗi xảy ra khi tạo sự kiện');
    }
  };

  return (
    // Your form JSX here
  );
};
```

### 3. Hiển thị danh sách sự kiện từ blockchain
```tsx
// src/pages/Index.tsx - cập nhật để lấy data từ blockchain
import { blockchainApi } from '@/services/blockchainApi';

const Index = () => {
  const [events, setEvents] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchEvents = async () => {
      try {
        const result = await blockchainApi.getContests();
        if (result.success) {
          setEvents(result.data);
        }
      } catch (error) {
        console.error('Failed to fetch events from blockchain:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchEvents();
  }, []);

  // Render events from blockchain data
  return (
    <div>
      {events.map(event => (
        <EventCard 
          key={event.id}
          title={event.name}
          description={event.description}
          date={event.start_date}
          verified={event.tx_hash ? true : false}
          txHash={event.tx_hash}
        />
      ))}
    </div>
  );
};
```

## 🔧 Cấu hình Blockchain thật

Để kết nối với blockchain thật:

1. **Deploy smart contract** (sử dụng contract hiện có của bạn)
2. **Cập nhật .env**:
   ```env
   CONTRACT_ADDRESS=0xYourContractAddress
   PRIVATE_KEY=0xYourPrivateKey
   NETWORK_URL=https://polygon-rpc.com
   ```

## 🏗️ Phát triển tiếp

- [ ] Implement ABI binding cho smart contract
- [ ] Thêm authentication/authorization
- [ ] Cache layer với Redis
- [ ] Database integration
- [ ] Rate limiting
- [ ] Logging và monitoring
- [ ] Docker containerization

## 🔒 Bảo mật

⚠️ **Quan trọng**: 
- Không commit private key thật vào git
- Sử dụng environment variables cho sensitive data
- Validate tất cả input từ user
- Implement rate limiting cho production

## 📝 License

MIT License - Sử dụng tự do cho mục đích học tập và phát triển.
