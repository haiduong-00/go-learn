package main

import "fmt"

func main() {
	const sotien_1_khach_chi = 120
	S := 0
	for i := 9; i < 21; i++ {
		if (i >= 9) && (i < 11) {
			S += 30 * sotien_1_khach_chi
		} else if (i >= 11) && (i < 13) {
			S += 120 * sotien_1_khach_chi
		} else if (i >= 13) && (i < 18) {
			S += 40 * sotien_1_khach_chi
		} else if (i >= 18) && (i < 21) {
			S += 150 * sotien_1_khach_chi
		}
	}
	fmt.Println("Tổng doanh thu 1 ngày của nhà hàng là:", S, "nghìn vnđ")
}
