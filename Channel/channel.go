package main

import (
	"fmt"
)

// Channel trong Concurrency
// Channel: như một dây chuyền sản xuất
// Có routine sản xuất, và có routine nhận (lấy) hàng ra
// Các routine chỉ thực hiện 1 hướng: routine nào lấy hàng thì sẽ chỉ lấy, routine sản xuất thì chỉ sản xuất (câu này sai)*
// Những routine phụ có thể lấy, hoặc sản suất. Nhưng riêng routine main thì chỉ có NHẬN
// Code bên dưới là 1 ví dụ về việc routine main không thể sản xuất
// Khi sản xuất muốn thu hồi lại sản phẩm: buffered channel: queue
// *Vì nếu có công việc thay ca, và sắp xếp việc sao cho channel có hàng: routine có thể thay ca nhau
// Muốn strict ca: routine này chỉ làm việc này, routine kia chỉ làm việc kia: Directional Channel: Send (sản xuất), Receive(Lấy, Nhận)
// Channel bình thường (General) thì có thể chuyển qua Directional Channel
// Directional Channel thì KHÔNG thể chuyển về dạng Channel bình thường (General)

func main() {
	c := make(chan int)
	go SendOnly(c)
	Receive(c)
}

func SendOnly(c chan<- int) {
	c <- 42
}

func Receive(c <-chan int) {
	fmt.Println(<-c)
}
