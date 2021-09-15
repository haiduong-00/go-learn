package main

import (
	"fmt"
)

// Đã từng học về Switch: tùy thuộc vào điều kiện của các case: mà sẽ chạy câu lệnh đó
// Select khá là tương tự, nhưng có 1 đặc điểm là dùng cho kiểm tra channel
// Select sẽ lọc ra sản phẩm đó đến từ channel nào

// Ok: check for channel: check xem channel đã bị đóng chưa

func main() {
	c := make(chan int)
	go func ()  {
		c <- 42
	}()
	v,ok:= <-c
	fmt.Println(v,ok)
	close(c)  // Đóng channel
	v,ok= <-c
	fmt.Println(v,ok)
}
