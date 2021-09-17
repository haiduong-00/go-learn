package main

import (
	"fmt"
	"log"
)

// Practice: Thói quen tốt: Check lỗi liên tục (nhưng có vài trường hợp ngoại lệ)
// Do error là 1 TYPE

// Go hỗ trợ rất nhiều cách in ra lỗi (sắp xếp theo mức độ nghiêm trọng):
// fmt.Print(); log.Print(), fmt.Errorf() (log sẽ trả ra cả ngày giờ gặp lỗi cụ thể) (Dạng chỉ IN ra lỗi)
// log.Panic(), panic() (Gây ra panicking, ảnh hưởng đến function và caller gọi ra function đó, theo hiệu ứng Domino)
// log.Fatal()  (Buộc dừng process ngay lập tức)

func check(err error) {
	if err != nil {
		log.Fatal(err) // Đặc biệt log còn có thể tạo ra 1 file log riêng
	}
}

func main() {
	err := fmt.Errorf("một lỗi gì đấy")
	check(err)
	fmt.Println("Hello")
}
