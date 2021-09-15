package main

import (
	"fmt"
)

// Range: for
// For - range statement: cái này được sử dụng để ĐI QUA tất cả giá trị của Array, Slice, Map, String (Slice of Byte)
// Channel cũng có range: range qua từng giá trị trong Channel
// Cho đến khi nhận báo kết thúc
// Trong những cấu trúc: Slice, Array: thì có index: i
// Trong map: thì có key: k
// Trong Channel (Linked list): nó chỉ có điểm đầu điểm cuối của dãy, trong dãy đấy có các gói hàng (node), cái node này chứa: hàng (data), địa chỉ tới gói hàng tiếp theo (pointer to next node)
// Channel không có index || key => Range của channel chỉ cần 1 giá trị

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}
