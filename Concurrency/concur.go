package main

import (
	"fmt"
	"sync"
)

// // Concurrency vs Parallelism
// // Concurrency: Đồng thời: ám chỉ tính chất 2 hay nhiều Processes chạy cùng lúc, và chia sẻ nhau dữ liệu mẹ
// // Parallelism: Song song: ám chỉ sự kiện 2 hay nhiều Processes chạy song song nhau

// // Code có thể áp dụng concurrency
// // Dù có Concurrency nhưng có thể KHÔNG xảy ra Parellelism

// // Những yếu tố nổi bật quyết định đến sự xảy ra của Parallelism
// // GOARCH: vi xử lý (CPU Architecture) hệ nào
// // GOOS: Hệ điều hành như thế nào (Những OS người dùng cuối (Windows, macOS, iOS), đều cho phép Parallelism, OS dùng trong điều khiển tàu vũ trụ không cho phép Parallelism xảy ra)
// // NumCPU(): số CPU (logically) của hệ thống logically: ám chỉ những gì mà Processes, Users, nhìn thấy, cho rằng là vậy
// // Threads: số Threads (logically) mà ứng dụng (Program), hay cả CPU, Hệ thống (System) tạo ra (Program != Process)
// // Threads là đơn vị đo của Process
// // vd: 1 process có 1 thread thì nó chạy không song song
// // Thread của ứng dụng ~ số CPU logically: 3 dạng liên kết many-to-one, one-to-one, many-to-many
// // Dựa vào Scheduler
// // Threads được tạo bởi ứng dụng Golang, nói threads đó là go routine
// // NumGoroutine(): số routine mà được tạo ra để chạy Parallelism

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// 	"sync/atomic"
// 	"time"
// )

// var wg sync.WaitGroup

// func main() {
// 	wg.Add(100)
// 	var con int64
// 	for i := 0; i < 100; i++ {
// 		go func() {

// 			// runtime.Gosched() // Cho phép những routine khác chạy
// 			// Data race
// 			// time.Sleep(time.Second * 2)
// 			atomic.AddInt64(&con, 1)
// 			fmt.Println(atomic.LoadInt64(&con))
// 			fmt.Println("routine:", runtime.NumGoroutine())
// 			fmt.Println("Heyyyyy")
// 			time.Sleep(time.Second * 2)
// 			fmt.Println("Heyyyyy")
// 			fmt.Println("Heyyyyy")
// 			fmt.Println("Heyyyyy")
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println(con)
// }

// Khắc phục data race trong Golang:
// trong package sync: synchronization: đồng bộ
// Mutex

// package sync/atomic
// Atomic:

// Một chương trình thông minh: là một chương trình có thể chạy parallelism hiệu quả nhất

func main()  {
	var wg sync.WaitGroup
	var num int
	wg.Add(1)
	go func ()  {
		num = foo()
		wg.Done()
	}() // Yêu cầu PHẢI là 1 function evoke, 1 method
	bar()
	wg.Wait()
	fmt.Println(num)
}

func foo() int {
	return 5
}

func bar() int {
	return 9
}
