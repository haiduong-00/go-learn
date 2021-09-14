package main

import (
	"fmt"
)

// Channel trong Concurrency
// Channel: như một dây chuyền sản xuất
// Có routine sản xuất, và có routine nhận (lấy) hàng ra
// Các routine chỉ thực hiện 1 hướng: routine nào lấy hàng thì sẽ chỉ lấy, routine sản xuất thì chỉ sản xuất
// Những routine phụ có thể lấy, hoặc sản suất. Nhưng riêng routine main thì chỉ có NHẬN
// Code bên dưới là 1 ví dụ về việc routine main không thể sản xuất

func main() {
	c := make(chan int) //
	quit := make(chan bool)
	c <- 42
	fmt.Println("Wait")
	go func() {
		fmt.Println(<-c)
		quit <- true
	}()
	<- quit 
}
