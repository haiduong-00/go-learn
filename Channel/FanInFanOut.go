package main

import (
	"fmt"
	"sync"
)

// Fan In, Fan Out: là một concept (mô hình) về cách xây dựng hệ thống concurrency
// Fan In: Lấy nhiều dữ liệu đưa về 1 luồng Output
// Fan Out: Phân dữ liệu thành nhiều luồng để xử lý

func FanInFanOut() {
	odd := make(chan int)
	even := make(chan int)
	c := make(chan int)
	var wg sync.WaitGroup
	go AddOdd(odd)
	go AddEven(even)
	go FanIn(odd, even, c)



	// Fan Out starting here
	for i := 0; i < 1000; i++ {
		//time.Sleep(time.Millisecond*5) // Mguy hiểm, gây giảm hiệu năng, giảm tốc độ
		wg.Add(1)
		go func() {  // Khi mà gặp câu lệnh go: chương trình sẽ nhảy đoạn code trong go đó ra routine khác, và vòng for tiếp tục công việc
			for v:=range c {
				if v%2 == 0 {
					fmt.Println("Even Number:", v)
				} else {
					fmt.Println("Odd Number: ", v)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()  // Hay sử dụng quit channel cũng tương đương: An toàn, không hao phí hiệu năng, khó kiểm soát, áp dụng
}

func AddOdd(odd chan<- int) {
	for i := 1; i <= 10000; i++ {
		if i%2 != 0 {
			odd <- i
		}
	}
	close(odd)
}
func AddEven(even chan<- int) {
	for i := 1; i <= 10000; i++ {
		if i%2 == 0 {
			even <- i
		}
	}
	close(even)
}

func FanIn(odd, even <-chan int, c chan int) {
	for {
		select {
		case v, ok := <-odd:
			if ok {
				c <- v
			} else {
				continue
			}
		case v, ok := <-even:
			if ok {
				c <- v
			} else {
				if _, ok := <-odd; !ok {
					close(c)
					return
				} else {
					continue
				}
			}
		}
	}
}
