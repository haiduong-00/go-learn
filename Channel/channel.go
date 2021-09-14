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

func main() {
	c := make(chan int) //
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		c <- 42
		fmt.Println(<-c)  
		wg.Done()
	}()
	go func() {
		c <- 38
		wg.Done()
	}()
	wg.Wait()
}


// Ví dụ nhỏ về channel:
// Login: nếu có xác thực 2 bước
// Tất cả những thứ sau đây đều chạy Parallelism:
// Trang web sẽ tạo code: sản xuất xong code rồi
// Trang web khởi động hệ thống kiểm tra code: code mẫu, chỉ chờ code từ người dùng
// Trang web tạo trang nhập code người dùng