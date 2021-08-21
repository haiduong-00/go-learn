// Tạo slice có giá trị sau {2015,2016,2017,2018,2019,2020,2021,2022,2023,2024,2025}
// Lấy và in ra những năm có Covid-19, chỉ dựa vào slice ban đầu
// Cắt bỏ những năm bị Covid-19 đi trong slice ban đầu. In ra slice đã cắt

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	x := []int{2015, 2016, 2017, 2018, 2019, 2020, 2021, 2022, 2023, 2024, 2025}
// 	fmt.Println(x[4:7])
// 	x = append(x[:4], x[8:]...)
// 	fmt.Println(x)
// 	fmt.Println(x)
// }

// Tạo slice có giá trị sau {2015,2016,2017,2018,2019,2020,2021,2022,2023,2024,2025}
// Lấy vào 1 slice mới và in ra những năm có Covid-19, chỉ dựa vào slice ban đầu
// Cắt bỏ những năm bị Covid-19 đi trong slice ban đầu. In ra slice đã cắt
// In ra x một lần nữa, nếu x đúng với slice đã cắt thì đã làm đúng

package main

import (
	"fmt"
)

func main() {
	x := []int{2015, 2016, 2017, 2018, 2019, 2020, 2021, 2022, 2023, 2024, 2025}
	xint := x[4:7]
	fmt.Println(xint)
	x = append(x[:4], x[7:]...)
	fmt.Println(x)
	fmt.Println(x)


}
