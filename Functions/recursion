package main

import (
	"fmt"
	"time"
)

func factorial(n int) int { //giai thừa
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// func timeToMeet(s float64, v float64) {
// 	V := 3 * v / 2
// 	time := s / V
// 	fmt.Println(time)
// }

// Recursion

// Prime number: số nguyên tố
// GCD, LCM: ước chung lớn nhất, bội chung nhỏ nhất
// Modular: % mod:
// Divisibility: tính chia hết
// Number Digits: bộ số 2, bộ số 16
// Triagle: tam giác
// Algebra: hàm số
//
// 5! = 5*4*3*2*1 = 4!
// n= 5: f(5) = 5*f(4)
// f(4) = 4*f(3)
// f(3) = 3* f(2)
// f(2) = 2*f(1)
// f(1) = 1 * f(0)
// f(0) = 0 * f(-1)

func main() {

	// fmt.Println(factorial(5))

	// Cái đệ quy (recursion) đã cũ lắm rồi
	start := time.Now()
	fac := 1
	for i := 1; i <= 5; i++ {
		fac *= i
	}
	fmt.Println(fac)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	
}


