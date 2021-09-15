package main

import "fmt"

// Đã từng học về Switch: tùy thuộc vào điều kiện của các case: mà sẽ chạy câu lệnh đó

func main() {
	n := 3
	switch n {
	case 3:
		fmt.Println("Hello")
	case 4:
		fmt.Println("Wrong Answer")
	}
}
