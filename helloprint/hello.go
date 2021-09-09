package hello

// Bước 1: Viết code
// Có thể liên kết nhiều code trong cùng 1 thư mục bằng cách sử dụng cùng tên package

// Bước 2: cd đến chỗ thư mục cần build cái package đấy
// ấn go mod init: Khởi tạo go.mod, đánh dấu địa chỉ package, đánh dấu go version, khởi tạo hệ thống quản lý module
// go mod tidy: đánh dấu những cái package bên ngoài, lưu vào go.sum, xóa dấu những cái package không cần trong go.sum, go.mod
// go list -m all: liệt kê những package bên ngoài, cả package đang xây dựng vào

// Bước 3: go build: xây dựng lên cái package đó cho sử dụng lâu dài

// go get

// Bước 4: Ở file go khác: import tên module trong go.mod

func Hello(name string) string {
	return print_hello(name)
}