package main

import (
	"fmt"
	"sync"
	"time"
)

// Channel trong Concurrency
// Channel: như một dây chuyền sản xuất
// Có routine sản xuất, và có routine nhận (lấy) hàng ra
// Các routine chỉ thực hiện 1 hướng: routine nào lấy hàng thì sẽ chỉ lấy, routine sản xuất thì chỉ sản xuất (câu này sai)*
// Những routine phụ có thể lấy, hoặc sản suất. Nhưng riêng routine main thì chỉ có NHẬN
// Code bên dưới là 1 ví dụ về việc routine main không thể sản xuất
// Khi sản xuất muốn thu hồi lại sản phẩm: buffered channel: queue
// *Vì nếu có công việc thay ca, và sắp xếp việc sao cho channel có hàng: routine có thể thay ca nhau

func main() {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		c <- 411
		time.Sleep(time.Second*2)
		fmt.Println(<-c)
		wg.Done()
	}()
	go func() {
		fmt.Println(<-c)
		c <- 55
		wg.Done()
	}()
	go func() {
		time.Sleep(time.Second)
		fmt.Println(<-c)
		c<-32
		wg.Done()
	}()
	wg.Wait()
}
