//Nguyên liệu: slice lưu các sự kiện, for i:=0; i<=91; switch case các sự kiện,
// if để xác định thời điểm in ra tỉ lệ (>phút thứ 40)

package main

import "fmt"

func main() {
	// Cơ hội: MCI: 10, TOT, -10
	// Sút trúng đích, cản phá: MCI: 2, TOT: -2
	// Thay người: MCI: 8, TOT -8
	// Thẻ vàng: MCI: -5, TOT: 5
	// Bàn thắng: MCI: 30, MCI: -30

	timeline := make([]int, 92)
	// timeline[5],timeline[35],timeline[70] = "CI","CI","CI"
	// timeline[6],timeline[40],timeline[60] = "CT","CT","CT"
	// timeline[24],timeline[75],timeline[84]= "ST","SI","SI"
	// timeline[70] = "CITI"
	// timeline[77] = crypto: bitcoin,

	for i := 0; i <= 91; i++ {
		switch i {
		case 5, 35:
			timeline[i] += 10
		case 6, 40, 60:
			timeline[i] -= 10
		case 24:
			timeline[i] -= 2
		case 75, 84:
			timeline[i] += 2
		case 70:
			timeline[i] += 8 + 10
		case 79:
			timeline[i] += 8 * 2
		case 77, 83, 90:
			timeline[i] -= 8
		case 66:
			timeline[i] += 5
		case 55:
			timeline[i] -= 30
		}
	}
	rate := 50;
	for i:=0;i<=91;i++{
		rate += timeline[i]
		switch i{
		case 40,77,91: fmt.Println(rate)
		}
	}
}
