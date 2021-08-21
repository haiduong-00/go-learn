// Tạo slice 2 chiều 2x2, gán trước giá trị tùy thích
// Thêm hàng dọc sau: {4,5} vào slice đó, in ra vị trí [2][0]   3x2 (hàng dọc x hàng ngang)
// Thêm các vào hàng ngang giá trị 10, in ra độ dài slice 3x3

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	x := [][]int{
// 		{1, 2},
// 		{3, 4},
// 	}
// 	x = append(x, [][]int{
// 		{4, 5},
// 	}...)
// 	fmt.Println(x[2][0])
// 	for i,v :=range x{  // v: copy của x[i]
// 		x[i] = append(v,10)
// 	}
// 	fmt.Println(len(x[0]))  //3x3
// }

// Tạo slice 2 chiều 2x2, gán trước giá trị tùy thích
// Thêm hàng dọc sau: {4,5} vào slice đó, in ra vị trí [2][0]   3x2 (hàng dọc x hàng ngang)
// Thêm các vào hàng ngang giá trị 10, in ra độ dài slice 3x3

package main

import "fmt"

func main() {
	x := [][]int{
		{1, 2},
		{3, 4},
	}
	x = append(x, [][]int{
		{4, 5},
	}...)
	fmt.Println(x[2][0])
	for i := range x {
		x[i] = append(x[i], 10)

	}
	fmt.Println(len(x))
}
