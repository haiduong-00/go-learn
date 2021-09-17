package main

import (
	"fmt"
	"time"
)

// What is Error:
// Sắp xếp theo mức độ khó chữa (theo cá nhân):

// 1: Không có error

// 2: Error Syntax
// 2: Undefined

// 3: Invalid Operations
// 3: Exceed limit (quá giới hạn (array[5] mà lại thêm 6 số))

// 4: Deadlock
// 4: Scope,....

// 5: Tổng hợp từ 2 trở lên trong 4 mức độ trên

// Error là lỗi, có thể gây ra do sử dụng sai cách, không phù hợp với ngôn ngữ, hay vượt quá giới hạn, ...
// Error thường làm một đoạn code bất hiệu
// Error Handling: Kiểm soát hệ thống nếu error xảy ra
// Trong Go: gộp chung exception và error vào 1 type: error

// Runtime error: trong 1 server chạy liên tục (runtime), có error sẽ khó bắt hơn vì
// Compiled-time error: bắt đầu build thì có lỗi

// Go is all about the TYPE

// Thường thì các ngôn ngữ hiện đại sẽ bắt exception bằng hệ thống try-catch-finally:
// try: thử đoạn code nhỏ xem có lỗi không vd:
// try {
// 	console.log(0/0)
// }
// catch: nếu có lỗi xảy ra thì báo hiệu và dừng (nếu không có finally)
// finally: đoạn code cuối cùng khi catch báo hiệu có lỗi
// a = 9 mà gặp lỗi: sử dụng finally để thoát a an toàn (có thể là a = 0)

// Go approach: sẽ kiểm tra lỗi mỗi khi 1 function hoạt động


type loi struct {
	name string
	when time.Time
}

func (t loi) Error() string {
	return fmt.Sprintln("Hello World!")
}

func main() {
	err := loi{"What",time.Now()}
	fmt.Printf("%T %v",err,err)
}