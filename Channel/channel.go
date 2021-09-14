package main

import (
	"fmt"
)

// Channel trong Concurrency
// Channel: như một dây chuyền sản xuất
// Có routine sản xuất, và có routine nhận (lấy) hàng ra
// Các routine chỉ thực hiện 1 hướng: routine nào lấy hàng thì sẽ chỉ lấy, routine sản xuất thì chỉ sản xuất

func main() {
	c := make(chan int) //
	quit := make(chan bool)
	go func() {
		c <- 32
		c <- 48
		c <- 56
	}()
	go func() {
		fmt.Println(<-c)
		fmt.Println(<-c)
		fmt.Println(<-c)
		quit <- true
	}()
	<- quit 
}
