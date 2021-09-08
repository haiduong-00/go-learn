package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//crypt: crypto, encrypt
	// Mã hóa
	// Cách mã hóa 1 mật khẩu:
	// Người dùng tạo tài khoản/mật khẩu: Mật khẩu (string)
	// Trang web, ứng dụng: mã hóa cái mật khẩu đấy và lưu trên hệ thống (bcrypt)
	// Khi mã hóa người ta sử dụng mã hóa 1 chiều (không có cách giải mã)
	// Khi người dùng đăng nhập: Mật khẩu (string)
	// Cái mật khẩu dùng đăng nhập: được mã hóa cùng kiểu (cùng cost, cùng salt) với mật khẩu trước (bcrypt)
	// Hệ thống so sánh 2 cái mã bcrypt với nhau

	Ps := []byte(`Hello`)
	Pscrypt, err := bcrypt.GenerateFromPassword(Ps, 12)
	check(err)
	fmt.Println(string(Pscrypt))

	Login := []byte(`Hello`)
	// Logincrypt,err := bcrypt.GenerateFromPassword(Login,12)
	check(err)
	err = bcrypt.CompareHashAndPassword(Pscrypt,Login)
	check(err)
	if err == nil {fmt.Println("Login accepted")}
	// Lưu mật khẩu trên chrome: có thể là string bình thường, Chrome sử dụng mã hóa 2 chiều (2 chiều lưu key mã hóa)

	// Nguyên tắc khi xây dựng trang đăng nhập: không bao giờ lưu mật khẩu người dùng
	// Tên tài khoản, email, sđt:
}
