# Kiểm thử hiệu năng với k6 cho Blockchain Demo API

Thư mục này chứa các kịch bản kiểm thử hiệu năng (performance testing) cho Blockchain Demo API sử dụng công cụ k6.

## Cài đặt k6

Để chạy các test hiệu năng, bạn cần cài đặt k6:

### Windows
```
winget install k6
```
hoặc
```
choco install k6
```

### macOS
```
brew install k6
```

### Linux
```
sudo apt-get install k6
```

Xem thêm hướng dẫn cài đặt tại: https://k6.io/docs/getting-started/installation/

## Cấu trúc test

Thư mục này chứa các kịch bản test khác nhau:

1. **k6-benchmark.js**: Test hiệu năng tổng quát cho tất cả các API endpoints
2. **k6-spike-test.js**: Test khả năng chịu tải đột biến (spike test)
3. **k6-soak-test.js**: Test khả năng chịu tải trong thời gian dài (soak test)

## Chạy test

### Điều kiện tiên quyết

1. Đảm bảo server API đang chạy (mặc định tại `http://localhost:8080`)
2. Đảm bảo blockchain service đã được khởi tạo và kết nối đúng

### Chạy test hiệu năng tổng quát

```
k6 run tests/k6-benchmark.js
```

### Chạy test khả năng chịu tải đột biến

```
k6 run tests/k6-spike-test.js
```

### Chạy test khả năng chịu tải trong thời gian dài

```
k6 run tests/k6-soak-test.js
```

### Tùy chỉnh thông số

Bạn có thể tùy chỉnh các thông số khi chạy test:

```
k6 run --vus 20 --duration 2m tests/k6-benchmark.js
```

- `--vus`: Số lượng người dùng ảo (virtual users)
- `--duration`: Thời gian chạy test

### Xuất kết quả

Xuất kết quả ra file JSON:

```
k6 run --out json=results.json tests/k6-benchmark.js
```

## Phân tích kết quả

Sau khi chạy test, k6 sẽ hiển thị các thông số quan trọng:

- **http_req_duration**: Thời gian phản hồi của các request
- **http_req_failed**: Tỷ lệ request thất bại
- **iterations**: Số lần lặp lại kịch bản test
- **vus**: Số lượng người dùng ảo
- **data_received/data_sent**: Lượng dữ liệu nhận/gửi

Các metrics tùy chỉnh:
- **error_rate**: Tỷ lệ lỗi
- **success_rate**: Tỷ lệ thành công
- **create_content_duration**: Thời gian tạo content
- **get_content_duration**: Thời gian lấy content
- **list_contents_duration**: Thời gian liệt kê contents
- **create_contest_duration**: Thời gian tạo contest
- **get_contest_duration**: Thời gian lấy contest
- **list_contests_duration**: Thời gian liệt kê contests
- **search_contests_duration**: Thời gian tìm kiếm contests
- **health_check_duration**: Thời gian kiểm tra health

## Tích hợp với CI/CD

Bạn có thể tích hợp các test này vào quy trình CI/CD bằng cách thêm các bước chạy k6 vào file cấu hình CI (như GitHub Actions, GitLab CI, Jenkins, v.v.).

Ví dụ với GitHub Actions:

```yaml
name: Performance Testing

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  k6_load_test:
    name: k6 Load Test
    runs-on: ubuntu-latest

    steps:
      - name: Checkout
        uses: actions/checkout@v2

      - name: Start API Server
        run: |
          # Khởi động server API của bạn
          # ...

      - name: Run k6 test
        uses: grafana/k6-action@v0.2.0
        with:
          filename: tests/k6-benchmark.js
```

