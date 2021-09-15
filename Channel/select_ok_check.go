package main

import (
	"fmt"
)

// Đã từng học về Switch: tùy thuộc vào điều kiện của các case: mà sẽ chạy câu lệnh đó
// Select khá là tương tự, nhưng có 1 đặc điểm là dùng cho kiểm tra channel
// Select sẽ lọc ra sản phẩm đó đến từ channel nào

// Ok: check for channel: check xem channel đã bị đóng chưa

func main() {
	odd := make(chan int)
	even := make(chan int)
	AddOddandEven(odd, even)
	for {
		select {
		case v, ok := <-odd:
			if ok {
				fmt.Println("This is odd number:", v)
			} else {
				continue
			}
		case v, ok := <-even:
			if ok {
				fmt.Println("This is even number:\t", v)
			} else {
				if _, ok := <-odd; !ok {
					return
				} else {
					continue
				}
			}
		}
	}
}

func AddOddandEven(odd, even chan<- int) {
	go func() {
		for i := 1; i < 101; i++ {
			if i%2 == 0 {
				even <- i
			} else {
				odd <- i
			}
		}
		close(odd)
		close(even)
	}()
}
