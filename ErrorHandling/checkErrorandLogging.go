package main

import "fmt"

// Practice: Thói quen tốt: Check lỗi liên tục (nhưng có vài trường hợp ngoại lệ)
// Do error là 1 TYPE

// Go hỗ trợ rất nhiều cách in ra lỗi:
// fmt.Print(); log.Print() (log sẽ trả ra cả ngày giờ gặp lỗi cụ thể) (Dạng chỉ IN ra lỗi)
// log.Panic(), panic() (Gây ra panicking, ảnh hưởng đến function và caller gọi ra function đó, theo hiệu ứng Domino)
// log.Fatal(),exit() (Buộc dừng process ngay lập tức)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	var y int
	_, err := fmt.Scanln(&y) // Nếu y nhập vào không là số
	check(err)
	fmt.Print(y) // Lệnh in này thường không cần check lỗi, nó chỉ có in giá trị bình thường ra
}
