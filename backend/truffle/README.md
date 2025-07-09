# ContentStorage Smart Contract

## Mục lục
- [Giới thiệu](#giới-thiệu)
- [Cấu trúc dự án](#cấu-trúc-dự-án)
- [Chức năng chính](#chức-năng-chính)
- [Triển khai & Kiểm thử](#triển-khai--kiểm-thử)
- [Liên hệ](#liên-hệ)

---

## Giới thiệu

`ContentStorage` là smart contract được phát triển bằng Solidity, phục vụ cho việc quản lý các cuộc thi, thí sinh, nhà tài trợ và nội dung wiki trên blockchain. Hợp đồng này hỗ trợ lưu trữ, truy xuất, tìm kiếm và xác thực dữ liệu một cách minh bạch, phi tập trung.

## Cấu trúc dự án

- `contracts/ContentStorage.sol`: Smart contract chính.
- `migrations/`: Script triển khai hợp đồng.
- `test/`: Các file kiểm thử hợp đồng.
- `build/`: Kết quả biên dịch hợp đồng.
- `scripts/`: Script hỗ trợ (ví dụ: cập nhật cấu hình).

## Chức năng chính

### 1. Quản lý nội dung wiki
- Thêm nội dung mới (`storeContent`)
- Lấy nội dung theo ID (`getContent`)
- Lấy danh sách tất cả nội dung (`getAllContentIds`)

### 2. Quản lý thí sinh
- Thêm thí sinh mới (`addContestant`)
- Lấy thông tin thí sinh (`getContestant`)
- Lấy danh sách thí sinh (`getAllContestantIds`)

### 3. Quản lý cuộc thi
- Tạo cuộc thi mới (`createContest`)
- Lưu contest dạng JSON (`createContestJson`)
- Lấy thông tin cuộc thi (`getContest`)
- Tìm kiếm cuộc thi theo từ khóa (`searchContests`)
- Lấy danh sách cuộc thi (`getAllContestIds`)

### 4. Quản lý nhà tài trợ
- Thêm nhà tài trợ (`addSponsor`)
- Lấy thông tin nhà tài trợ (`getSponsor`)
- Lấy danh sách nhà tài trợ (`getAllSponsorIds`)

### 5. Đăng ký thí sinh vào cuộc thi
- Đăng ký thí sinh (`registerContestant`)
- Kiểm tra đăng ký (`isContestantRegistered`)
- Lấy danh sách thí sinh đã đăng ký một cuộc thi (`getContestantsInContest`)

## Triển khai & Kiểm thử

1. **Cài đặt dependencies:**
   ```bash
   cd backend/truffle
   npm install
   ```
2. **Biên dịch hợp đồng:**
   ```bash
   npx truffle compile
   ```
3. **Triển khai hợp đồng lên local blockchain:**
   ```bash
   npx truffle migrate
   ```
4. **Chạy kiểm thử:**
   ```bash
   npx truffle test
   ```


