
# Dự án Blockchain Contest & Wiki Platform

>Một hệ thống quản lý cuộc thi, thí sinh, nhà tài trợ và nội dung wiki dựa trên blockchain, sử dụng smart contract Solidity, backend Go, và frontend hiện đại với Vite + React + TailwindCSS.

---

## Mục lục
- [Giới thiệu](#giới-thiệu)
- [Kiến trúc tổng thể](#kiến-trúc-tổng-thể)
- [Các thành phần chính](#các-thành-phần-chính)
- [Hướng dẫn cài đặt & chạy thử](#hướng-dẫn-cài-đặt--chạy-thử)
- [Smart Contract: ContentStorage](#smart-contract-contentstorage)
- [Backend Go API](#backend-go-api)
- [Frontend Globe Wiki Canvas](#frontend-globe-wiki-canvas)
- [Đóng góp & phát triển](#đóng-góp--phát-triển)
- [Liên hệ](#liên-hệ)

---

## Giới thiệu

Dự án này xây dựng một nền tảng quản lý cuộc thi, thí sinh, nhà tài trợ và nội dung wiki minh bạch, phi tập trung trên blockchain. Hệ thống gồm 3 phần:

- **Smart contract Solidity**: Lưu trữ, xác thực và truy vấn dữ liệu trên blockchain.
- **Backend Go**: API trung gian, tích hợp blockchain và các dịch vụ khác.
- **Frontend React**: Giao diện người dùng hiện đại, trực quan.

## Kiến trúc tổng thể

```
┌────────────┐      ┌──────────────┐      ┌──────────────┐
│  Frontend  │ <--> │   Backend    │ <--> │  Blockchain  │
│ (React)    │      │   (Go API)   │      │ (Solidity)   │
└────────────┘      └──────────────┘      └──────────────┘
```

## Các thành phần chính

### 1. Smart Contract (Solidity)
- Quản lý cuộc thi, thí sinh, nhà tài trợ, nội dung wiki.
- Đảm bảo minh bạch, không thể sửa/xóa dữ liệu đã ghi.
- Hỗ trợ tìm kiếm, xác thực, đăng ký tham gia cuộc thi.

### 2. Backend Go
- API RESTful kết nối frontend và blockchain.
- Xử lý logic nghiệp vụ, xác thực, tích hợp các dịch vụ ngoài.

### 3. Frontend (Vite + React + TailwindCSS)
- Giao diện người dùng hiện đại, responsive.
- Cho phép tạo, xem, tìm kiếm cuộc thi, thí sinh, nội dung wiki.

## Hướng dẫn cài đặt & chạy thử

### 1. Smart Contract (Truffle)

```bash
cd backend/truffle
npm install
npx truffle compile
npx truffle migrate
npx truffle test
```

### 2. Backend Go

```bash
cd blockchain-demo-go
go mod tidy
go run cmd/server/main.go
```

### 3. Frontend (Vite + React)

```bash
cd globe-wiki-canvas
npm install
npm run dev
```

## Smart Contract: ContentStorage

File: `backend/truffle/contracts/ContentStorage.sol`

### Chức năng chính:
- **Quản lý nội dung wiki**: Thêm, lấy, liệt kê nội dung.
- **Quản lý thí sinh**: Thêm, lấy, liệt kê thí sinh.
- **Quản lý cuộc thi**: Tạo, lấy, tìm kiếm, lưu contest dạng JSON.
- **Quản lý nhà tài trợ**: Thêm, lấy, liệt kê nhà tài trợ.
- **Đăng ký thí sinh vào cuộc thi**: Đăng ký, kiểm tra, liệt kê thí sinh đã đăng ký.

### Một số hàm tiêu biểu:
- `storeContent`, `getContent`, `getAllContentIds`
- `addContestant`, `getContestant`, `getAllContestantIds`
- `createContest`, `getContest`, `searchContests`, `getAllContestIds`, `createContestJson`
- `addSponsor`, `getSponsor`, `getAllSponsorIds`
- `registerContestant`, `isContestantRegistered`, `getContestantsInContest`

## Backend Go API

Thư mục: `blockchain-demo-go/`

- Kết nối smart contract, cung cấp API cho frontend.
- Xử lý xác thực, kiểm thử, tích hợp các dịch vụ ngoài.
- Có thể mở rộng cho các nghiệp vụ khác.

## Frontend Globe Wiki Canvas

Thư mục: `globe-wiki-canvas/`

- Ứng dụng web hiện đại, sử dụng Vite, React, TailwindCSS.
- Cho phép người dùng tạo, xem, tìm kiếm các cuộc thi, thí sinh, nội dung wiki.
- Tích hợp với backend Go và smart contract.

## Đóng góp & phát triển

Mọi đóng góp đều được hoan nghênh! Vui lòng tạo pull request hoặc issue nếu bạn muốn tham gia phát triển hoặc báo lỗi.

## Liên hệ

- Tác giả: [Tên của bạn]
- Email: [Email của bạn]
- Github: [Link Github của bạn]

---
*Vui lòng thay thế thông tin liên hệ bằng thông tin của bạn.*
