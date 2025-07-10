// Cấu hình chung cho các test k6
export const BASE_URL = 'http://localhost:8080/api/v1';

// Hàm thiết lập headers
export function getHeaders() {
  return {
    'Content-Type': 'application/json',
  };
}

// Cấu hình thresholds (ngưỡng) mặc định
export const defaultThresholds = {
  'http_req_duration': ['p(95)<2000'],  // 95% requests hoàn thành trong 2s
  'error_rate': ['rate<0.1'],           // Tỷ lệ lỗi < 10%
};

// Cấu hình các giai đoạn mặc định cho load test
export const loadTestStages = [
  { duration: '30s', target: 10 },  // Tăng dần lên 10 users trong 30s
  { duration: '1m', target: 10 },   // Giữ ổn định 10 users trong 1 phút
  { duration: '20s', target: 0 },   // Giảm dần về 0 user trong 20s
];

// Cấu hình các giai đoạn cho spike test
export const spikeTestStages = [
  { duration: '10s', target: 5 },    // Khởi động với 5 users
  { duration: '20s', target: 5 },    // Giữ ổn định
  { duration: '10s', target: 50 },   // Tăng đột biến lên 50 users
  { duration: '30s', target: 50 },   // Duy trì tải cao
  { duration: '10s', target: 5 },    // Giảm về mức bình thường
  { duration: '10s', target: 0 },    // Kết thúc
];

// Cấu hình các giai đoạn cho soak test
export const soakTestStages = [
  { duration: '1m', target: 5 },     // Tăng dần lên 5 users
  { duration: '5m', target: 5 },     // Giữ ổn định 5 users trong 5 phút
  { duration: '30s', target: 0 },    // Giảm dần về 0
];

// Danh sách endpoints chính
export const endpoints = {
  health: `${BASE_URL}/health`,
  content: {
    create: `${BASE_URL}/content`,
    list: `${BASE_URL}/contents`,
    get: (id) => `${BASE_URL}/content/${id}`,
  },
  contest: {
    create: `${BASE_URL}/contests`,
    list: `${BASE_URL}/contests`,
    get: (id) => `${BASE_URL}/contests/${id}`,
    search: (query) => `${BASE_URL}/contests/search?q=${query}`,
  },
  contestant: {
    create: `${BASE_URL}/contestants`,
    list: `${BASE_URL}/contestants`,
    get: (id) => `${BASE_URL}/contestants/${id}`,
  },
  sponsor: {
    create: `${BASE_URL}/sponsors`,
    list: `${BASE_URL}/sponsors`,
    get: (id) => `${BASE_URL}/sponsors/${id}`,
  },
  registration: {
    register: (contestId) => `${BASE_URL}/contests/${contestId}/register`,
    getContestants: (contestId) => `${BASE_URL}/contests/${contestId}/contestants`,
  },
  stats: `${BASE_URL}/stats`,
};

// Hàm tạo payload mẫu cho content
export function createContentPayload(title = null, content = null, creator = null) {
  return {
    title: title || `Test Content ${Math.floor(Math.random() * 10000)}`,
    content: content || `This is test content created at ${new Date().toISOString()}`,
    creator: creator || 'k6-test'
  };
}

// Hàm tạo payload mẫu cho contest
export function createContestPayload(name = null, description = null) {
  const startDate = new Date();
  startDate.setDate(startDate.getDate() + 1);
  
  const endDate = new Date();
  endDate.setDate(endDate.getDate() + 7);
  
  return {
    name: name || `Test Contest ${Math.floor(Math.random() * 10000)}`,
    description: description || `This is a test contest created at ${new Date().toISOString()}`,
    start_date: startDate.toISOString(),
    end_date: endDate.toISOString(),
    image_url: 'https://example.com/test-image.jpg'
  };
}

// Hàm tạo payload mẫu cho contestant
export function createContestantPayload(name = null, details = null) {
  return {
    name: name || `Test Contestant ${Math.floor(Math.random() * 10000)}`,
    details: details || `This is a test contestant created at ${new Date().toISOString()}`,
    creator: 'k6-test'
  };
}

// Hàm tạo payload mẫu cho sponsor
export function createSponsorPayload(name = null, contactInfo = null) {
  return {
    name: name || `Test Sponsor ${Math.floor(Math.random() * 10000)}`,
    contact_info: contactInfo || `contact@sponsor${Math.floor(Math.random() * 10000)}.com`,
    sponsorship_amount: Math.floor(Math.random() * 10000) + 1000
  };
} 