package main

import (
	"fmt"
	"sync"
)

// Channel trong Concurrency
// Channel: như một dây chuyền sản xuất
// Có routine sản xuất, và có routine nhận (lấy) hàng ra
// Các routine chỉ thực hiện 1 hướng: routine nào lấy hàng thì sẽ chỉ lấy, routine sản xuất thì chỉ sản xuất
// Những routine phụ có thể lấy, hoặc sản suất. Nhưng riêng routine main thì chỉ có NHẬN
// Code bên dưới là 1 ví dụ về việc routine main không thể sản xuất
// Khi sản xuất muốn thu hồi lại sản phẩm: buffered channel: queue

func main() {
	c := make(chan int, 100)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		c <- 22
		wg.Done()
	}()
	go func() {
		c <- 42
		c <- 38
		fmt.Println(<-c)
		fmt.Println(<-c)
		fmt.Println(<-c)
		wg.Done()
	}()
	
	wg.Wait()
}
