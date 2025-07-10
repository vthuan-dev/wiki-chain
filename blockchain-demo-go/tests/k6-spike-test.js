import http from 'k6/http';
import { check, group, sleep } from 'k6';
import { Rate, Trend } from 'k6/metrics';
import { 
  getHeaders, 
  spikeTestStages, 
  endpoints,
  createContentPayload
} from './k6-config.js';

// Tùy chỉnh metrics
const errorRate = new Rate('error_rate');
const apiLatency = new Trend('api_latency');

// Cấu hình spike test - kiểm tra khả năng chịu tải đột biến
export const options = {
  stages: spikeTestStages,
  thresholds: {
    'http_req_duration': ['p(95)<5000'], // 95% requests phải hoàn thành trong 5s khi tải cao
    'error_rate': ['rate<0.2'],          // Tỷ lệ lỗi < 20% trong điều kiện tải cao
  },
};

// Hàm chính
export default function() {
  group('API Endpoints Under Spike Load', () => {
    // Test health check - endpoint nhẹ nhất
    {
      const startTime = new Date();
      const res = http.get(endpoints.health, { headers: getHeaders() });
      const duration = new Date() - startTime;
      
      apiLatency.add(duration, { endpoint: 'health' });
      check(res, {
        'health check is successful': (r) => r.status === 200,
      }) || errorRate.add(1);
    }
    
    // Test list contests - endpoint đọc dữ liệu
    {
      const startTime = new Date();
      const res = http.get(endpoints.contest.list, { headers: getHeaders() });
      const duration = new Date() - startTime;
      
      apiLatency.add(duration, { endpoint: 'list_contests' });
      check(res, {
        'list contests is successful': (r) => r.status === 200,
      }) || errorRate.add(1);
    }
    
    // Test search contests - endpoint tìm kiếm
    {
      const searchTerms = ['test', 'performance', 'contest'];
      const randomTerm = searchTerms[Math.floor(Math.random() * searchTerms.length)];
      
      const startTime = new Date();
      const res = http.get(endpoints.contest.search(randomTerm), { headers: getHeaders() });
      const duration = new Date() - startTime;
      
      apiLatency.add(duration, { endpoint: 'search_contests' });
      check(res, {
        'search contests is successful': (r) => r.status === 200,
      }) || errorRate.add(1);
    }
    
    // Test create content - endpoint ghi dữ liệu (nặng nhất)
    {
      const payload = JSON.stringify(createContentPayload(
        `Spike Test Content ${Math.floor(Math.random() * 1000)}`,
        `This is a spike test content created at ${new Date().toISOString()}`,
        'k6-spike-test'
      ));

      const startTime = new Date();
      const res = http.post(endpoints.content.create, payload, { headers: getHeaders() });
      const duration = new Date() - startTime;
      
      apiLatency.add(duration, { endpoint: 'create_content' });
      check(res, {
        'create content is successful': (r) => r.status === 201,
      }) || errorRate.add(1);
    }
  });
  
  // Nghỉ ngắn giữa các vòng lặp
  sleep(1);
} 