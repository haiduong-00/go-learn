package main

import (
	"fmt"
	// "math"
)


func main() {
	// Function sẽ được đánh giá, được coi như 1 value, thì có thể gán cho 1 variable,.....
	// f := func(x int) float64 {  // trở thành 1 expression để gán cho identifier
	// 	return math.Pow(float64(x),2)+float64(x)+1 // yêu cầu 1 expression, trả về kết quả là 1 func?
	// }

	// Callback & Closure
	// Rất có lợi cho closure
	// Closure: đơn giản là limit the scope of variable (giới hạn cái scope của variable)
	// Giảm sức nặng bộ nhớ
	bar := addN(3)  // func env: sẽ chứa n 
	// Recursion
	//  bar = func(x int) int {
	// 		return x+n
	// 		}	
	fmt.Println(bar(5))

	// // Dựa vào ý đồ mình xây dựng func để làm gì: return sẽ có giá trị trả về (kết quả) tùy thuộc
	// opDice := func (n int) int { // gọi ra cái int để chứng minh func expression có thể ở rất nhiều loại: có thể có return hoặc không, có thể có para hoặc không
	// 	return 6 - n + 1
	// }

	// y:=f(3) // y = f(3) = 3^2 + 3 + 1 = 13
	// fmt.Println(y)
	// fmt.Println(opDice(5))


	// f(x) = x^2 + x +1 = f(2) = return 
	// 1 nguyên tắc sử dụng return: phải liên quan đến các câu lệnh trong hàm số (func) đó
}

func addN(NumIn int) func (x int) int {  // NumIn= 5
	// n := NumIn  // n =5
	return ( func(x int) int {
		return x+NumIn
	}	)
} 

// func foo() int {
// 	fmt.Println("foo")
//  return 0
// }
