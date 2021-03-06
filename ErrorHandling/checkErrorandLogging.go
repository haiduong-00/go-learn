package main

import (
	"fmt"
)

// Practice: Thói quen tốt: Check lỗi liên tục (nhưng có vài trường hợp ngoại lệ)
// Do error là 1 TYPE

// Go hỗ trợ rất nhiều cách in ra lỗi (sắp xếp theo mức độ nghiêm trọng):
// fmt.Print(); log.Print(), fmt.Errorf() (log sẽ trả ra cả ngày giờ gặp lỗi cụ thể) (Dạng chỉ IN ra lỗi)
// log.Panic(), panic() (Gây ra panicking, ảnh hưởng đến function và caller gọi ra function đó, theo hiệu ứng Domino)
// log.Fatal()  (Buộc dừng process ngay lập tức)

// Chơi không lại, đái vào server (rút điện server)

func check(err error) {
	if err != nil {
		panic(err) // Đặc biệt log còn có thể tạo ra 1 file log riêng
	}
}

func chekError() {
	err := fmt.Errorf("một lỗi gì đấy")
	check(err)
	fmt.Println("Hello")
}
