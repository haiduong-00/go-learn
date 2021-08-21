// Tạo slice 2 chiều Oxy 3x3; gán số tùy thích, in ra số ở vị trí [1][2]

// package main
// import("fmt")
// func main() {
// 	x:=[][]int {
// 		{1,2,3},
// 		{4,5,6},
// 		{7,8,9},
// 	}
// 	xOneD := []int{3,4,5,6}
// 	fmt.Println(x[1][2])
// 	fmt.Println(xOneD[len(xOneD)-1])
// }

// Tạo slice 2 chiều Oxy 3x3; gán số tùy thích, in ra số ở vị trí [1][2]

package main

import (
	"fmt"
)

func main() {
	x := [][]int{
		{1, 2, 7},
		{3, 4, 8},
		{5, 6, 9},
	}
	fmt.Println(x[1][2])
}
