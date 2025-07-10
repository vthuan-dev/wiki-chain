import http from 'k6/http';
import { check, group, sleep } from 'k6';
import { Rate, Trend } from 'k6/metrics';
import { randomString } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

// Tùy chỉnh metrics
const errorRate = new Rate('error_rate');
const memoryLeakIndicator = new Trend('response_time_trend');

// Cấu hình soak test - kiểm tra khả năng chịu tải trong thời gian dài
export const options = {
  stages: [
    { duration: '1m', target: 5 },     // Tăng dần lên 5 users
    { duration: '5m', target: 5 },     // Giữ ổn định 5 users trong 5 phút
    { duration: '30s', target: 0 },    // Giảm dần về 0
  ],
  thresholds: {
    'http_req_duration': ['p(95)<3000'],           // 95% requests hoàn thành trong 3s
    'error_rate': ['rate<0.05'],                   // Tỷ lệ lỗi < 5%
    'response_time_trend': ['trend()<200'],        // Thời gian phản hồi không tăng quá 200ms
  },
};

// URL cơ sở của API
const BASE_URL = 'http://localhost:8080/api/v1';

// Hàm thiết lập headers
function getHeaders() {
  return {
    'Content-Type': 'application/json',
  };
}

// Biến lưu trữ dữ liệu giữa các requests
let contentIds = [];
let contestIds = [];

// Hàm chính
export default function() {
  // Theo dõi thời gian phản hồi để phát hiện memory leak
  const startTime = new Date();
  
  group('Long-Running API Tests', () => {
    // Kiểm tra health endpoint định kỳ
    {
      const res = http.get(`${BASE_URL}/health`, { headers: getHeaders() });
      check(res, {
        'health check remains healthy': (r) => r.status === 200 && r.json('status') === 'healthy',
      }) || errorRate.add(1);
    }
    
    // Tạo content mới định kỳ
    if (Math.random() < 0.3) { // 30% khả năng thực hiện
      const payload = JSON.stringify({
        title: `Soak Test Content ${randomString(8)}`,
        content: `This is a soak test content created at ${new Date().toISOString()}`,
        creator: 'k6-soak-test'
      });

      const res = http.post(`${BASE_URL}/content`, payload, { headers: getHeaders() });
      
      const success = check(res, {
        'create content remains successful': (r) => r.status === 201,
      });
      
      if (success && res.json('id')) {
        // Giới hạn số lượng ID lưu trữ để tránh memory leak trong k6
        if (contentIds.length > 50) {
          contentIds.shift(); // Loại bỏ ID cũ nhất
        }
        contentIds.push(res.json('id'));
      } else {
        errorRate.add(1);
      }
    }
    
    // Lấy danh sách content định kỳ
    if (Math.random() < 0.4) { // 40% khả năng thực hiện
      const res = http.get(`${BASE_URL}/contents`, { headers: getHeaders() });
      check(res, {
        'list contents remains successful': (r) => r.status === 200,
        'list contents returns data consistently': (r) => Array.isArray(r.json('data')),
      }) || errorRate.add(1);
    }
    
    // Lấy chi tiết content nếu có ID
    if (contentIds.length > 0 && Math.random() < 0.3) {
      const randomIndex = Math.floor(Math.random() * contentIds.length);
      const contentId = contentIds[randomIndex];
      
      const res = http.get(`${BASE_URL}/content/${contentId}`, { headers: getHeaders() });
      check(res, {
        'get content remains successful': (r) => r.status === 200,
      }) || errorRate.add(1);
    }
    
    // Tạo contest mới định kỳ
    if (Math.random() < 0.2) { // 20% khả năng thực hiện
      const startDate = new Date();
      startDate.setDate(startDate.getDate() + 1);
      
      const endDate = new Date();
      endDate.setDate(endDate.getDate() + 7);
      
      const payload = JSON.stringify({
        name: `Soak Test Contest ${randomString(8)}`,
        description: `This is a soak test contest created at ${new Date().toISOString()}`,
        start_date: startDate.toISOString(),
        end_date: endDate.toISOString(),
        image_url: 'https://example.com/test-image.jpg'
      });

      const res = http.post(`${BASE_URL}/contests`, payload, { headers: getHeaders() });
      
      const success = check(res, {
        'create contest remains successful': (r) => r.status === 201,
      });
      
      if (success && res.json('id')) {
        // Giới hạn số lượng ID lưu trữ
        if (contestIds.length > 30) {
          contestIds.shift();
        }
        contestIds.push(res.json('id'));
      } else {
        errorRate.add(1);
      }
    }
    
    // Lấy danh sách contest định kỳ
    if (Math.random() < 0.4) {
      const res = http.get(`${BASE_URL}/contests`, { headers: getHeaders() });
      check(res, {
        'list contests remains successful': (r) => r.status === 200,
      }) || errorRate.add(1);
    }
    
    // Tìm kiếm contest định kỳ
    if (Math.random() < 0.3) {
      const searchTerms = ['test', 'soak', 'performance'];
      const randomTerm = searchTerms[Math.floor(Math.random() * searchTerms.length)];
      
      const res = http.get(`${BASE_URL}/contests/search?q=${randomTerm}`, { headers: getHeaders() });
      check(res, {
        'search contests remains successful': (r) => r.status === 200,
      }) || errorRate.add(1);
    }
    
    // Kiểm tra thống kê định kỳ
    if (Math.random() < 0.2) {
      const res = http.get(`${BASE_URL}/stats`, { headers: getHeaders() });
      check(res, {
        'stats endpoint remains successful': (r) => r.status === 200,
      }) || errorRate.add(1);
    }
  });
  
  // Đo thời gian phản hồi tổng thể để phát hiện memory leak
  const duration = new Date() - startTime;
  memoryLeakIndicator.add(duration);
  
  // Nghỉ ngắn giữa các vòng lặp - thời gian ngẫu nhiên để mô phỏng người dùng thực
  sleep(Math.random() * 2 + 1); // 1-3 giây
} 