import http from 'k6/http';
import { check, group, sleep } from 'k6';
import { Counter, Rate, Trend } from 'k6/metrics';
import { randomString } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';
import { 
  BASE_URL, 
  getHeaders, 
  loadTestStages, 
  defaultThresholds,
  endpoints,
  createContentPayload,
  createContestPayload
} from './k6-config.js';

// Tùy chỉnh metrics
const errorRate = new Rate('error_rate');
const successRate = new Rate('success_rate');
const createContentTrend = new Trend('create_content_duration');
const getContentTrend = new Trend('get_content_duration');
const listContentsTrend = new Trend('list_contents_duration');
const createContestTrend = new Trend('create_contest_duration');
const getContestTrend = new Trend('get_contest_duration');
const listContestsTrend = new Trend('list_contests_duration');
const searchContestsTrend = new Trend('search_contests_duration');
const healthCheckTrend = new Trend('health_check_duration');

// Cấu hình mặc định
export const options = {
  stages: loadTestStages,
  thresholds: {
    ...defaultThresholds,
    'create_content_duration': ['p(95)<3000'], // 95% tạo content trong 3s
    'create_contest_duration': ['p(95)<3000'],  // 95% tạo contest trong 3s
    'health_check_duration': ['p(99)<500'],     // 99% health check trong 0.5s
  },
};

// Biến lưu trữ dữ liệu giữa các requests
let contentIds = [];
let contestIds = [];

// Hàm chính
export default function() {
  group('Health Check', () => {
    const startTime = new Date();
    const res = http.get(endpoints.health, { headers: getHeaders() });
    const duration = new Date() - startTime;
    
    healthCheckTrend.add(duration);
    check(res, {
      'health check status is 200': (r) => r.status === 200,
      'health check reports healthy': (r) => r.json('status') === 'healthy',
    }) || errorRate.add(1);
    successRate.add(res.status === 200 ? 1 : 0);
  });

  group('Content API', () => {
    // Tạo content mới
    {
      const payload = JSON.stringify(createContentPayload(
        `Performance Test Content ${randomString(8)}`,
        `This is a performance test content created by k6 at ${new Date().toISOString()}`,
        'k6-performance-test'
      ));

      const startTime = new Date();
      const res = http.post(endpoints.content.create, payload, { headers: getHeaders() });
      const duration = new Date() - startTime;
      
      createContentTrend.add(duration);
      const success = check(res, {
        'create content status is 201': (r) => r.status === 201,
        'create content response has ID': (r) => r.json('id') !== undefined,
      });
      
      if (success && res.json('id')) {
        contentIds.push(res.json('id'));
      } else {
        errorRate.add(1);
      }
      
      successRate.add(success ? 1 : 0);
      sleep(1);
    }

    // Lấy danh sách content
    {
      const startTime = new Date();
      const res = http.get(endpoints.content.list, { headers: getHeaders() });
      const duration = new Date() - startTime;
      
      listContentsTrend.add(duration);
      check(res, {
        'list contents status is 200': (r) => r.status === 200,
        'list contents returns array': (r) => Array.isArray(r.json('data')),
      }) || errorRate.add(1);
      
      successRate.add(res.status === 200 ? 1 : 0);
      sleep(0.5);
    }

    // Lấy chi tiết content nếu có ID
    if (contentIds.length > 0) {
      const randomIndex = Math.floor(Math.random() * contentIds.length);
      const contentId = contentIds[randomIndex];
      
      const startTime = new Date();
      const res = http.get(endpoints.content.get(contentId), { headers: getHeaders() });
      const duration = new Date() - startTime;
      
      getContentTrend.add(duration);
      check(res, {
        'get content status is 200': (r) => r.status === 200,
        'get content returns correct data': (r) => r.json('data.id') === contentId,
      }) || errorRate.add(1);
      
      successRate.add(res.status === 200 ? 1 : 0);
      sleep(0.5);
    }
  });

  group('Contest API', () => {
    // Tạo contest mới
    {
      const payload = JSON.stringify(createContestPayload(
        `Performance Test Contest ${randomString(8)}`,
        `This is a performance test contest created by k6 at ${new Date().toISOString()}`
      ));

      const startTime = new Date();
      const res = http.post(endpoints.contest.create, payload, { headers: getHeaders() });
      const duration = new Date() - startTime;
      
      createContestTrend.add(duration);
      const success = check(res, {
        'create contest status is 201': (r) => r.status === 201,
        'create contest response has ID': (r) => r.json('id') !== undefined,
      });
      
      if (success && res.json('id')) {
        contestIds.push(res.json('id'));
      } else {
        errorRate.add(1);
      }
      
      successRate.add(success ? 1 : 0);
      sleep(1);
    }

    // Lấy danh sách contest
    {
      const startTime = new Date();
      const res = http.get(endpoints.contest.list, { headers: getHeaders() });
      const duration = new Date() - startTime;
      
      listContestsTrend.add(duration);
      check(res, {
        'list contests status is 200': (r) => r.status === 200,
        'list contests returns array': (r) => Array.isArray(r.json('data')),
      }) || errorRate.add(1);
      
      successRate.add(res.status === 200 ? 1 : 0);
      sleep(0.5);
    }

    // Tìm kiếm contest
    {
      const searchTerms = ['test', 'performance', 'contest'];
      const randomTerm = searchTerms[Math.floor(Math.random() * searchTerms.length)];
      
      const startTime = new Date();
      const res = http.get(endpoints.contest.search(randomTerm), { headers: getHeaders() });
      const duration = new Date() - startTime;
      
      searchContestsTrend.add(duration);
      check(res, {
        'search contests status is 200': (r) => r.status === 200,
        'search contests returns data': (r) => r.json('data') !== undefined,
      }) || errorRate.add(1);
      
      successRate.add(res.status === 200 ? 1 : 0);
      sleep(0.5);
    }

    // Lấy chi tiết contest nếu có ID
    if (contestIds.length > 0) {
      const randomIndex = Math.floor(Math.random() * contestIds.length);
      const contestId = contestIds[randomIndex];
      
      const startTime = new Date();
      const res = http.get(endpoints.contest.get(contestId), { headers: getHeaders() });
      const duration = new Date() - startTime;
      
      getContestTrend.add(duration);
      check(res, {
        'get contest status is 200': (r) => r.status === 200,
        'get contest returns correct data': (r) => r.json('data.id') === contestId,
      }) || errorRate.add(1);
      
      successRate.add(res.status === 200 ? 1 : 0);
      sleep(0.5);
    }
  });

  // Kiểm tra thống kê
  group('Stats API', () => {
    const res = http.get(endpoints.stats, { headers: getHeaders() });
    
    check(res, {
      'stats status is 200': (r) => r.status === 200,
      'stats returns data': (r) => r.json('data') !== undefined,
    }) || errorRate.add(1);
    
    successRate.add(res.status === 200 ? 1 : 0);
  });

  // Nghỉ ngắn giữa các vòng lặp
  sleep(1);
} 