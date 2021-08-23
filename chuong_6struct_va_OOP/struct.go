package main

import "fmt"

// Non-anonymous struct

type person struct {
	height   int
	weight   int
	fullName string
}

type Duong struct {
	person  // Embedded type
	footballSkillRanked int
	footballPassion bool
}



// type <tên kiểu> struct {
// 		<tên đặc điểm> <kiểu dữ liệu>
// 		<tên đặc điểm> <kiểu dữ liệu>
// }


func (t *Duong) fastSprint()  {
	fmt.Println("Dương đang dốc bóc rất nhanh đến khu vực đội bạn")
}

func (t *person) eat()  {
	fmt.Printf("%v đang ăn", t.fullName)
}

func main() {
	Duong := Duong{ // Object: Duong
		person: person{
			height: 178,
			weight: 70,
			fullName: "Đinh Hải Dương",
		},
		footballSkillRanked: 8,
		footballPassion: true,
	}
	fmt.Println(Duong.fullName)

	Duong.fastSprint()
	Duong.eat()

	// <tên đối tượng> := <tên kiểu> {
	// 		<đặc điểm>: <giá trị>,
	// }

	// Để lấy ra đặc điểm, hành động: <tên đối tượng>.<dặc điểm>;  <tên đối tượng>.<hành động>(parameters)

	// Object: Đặc điểm (fields, state, data), Hành động (methods, behavior)
	// OOP: Object-orientation programming: Lập trình hướng đối tượng

	// map có ưu điểm: tìm thông tin rất nhanh
	// So với slice, array
	// Đôi khi vẫn sử dụng slice:
	// Sử dụng kết hợp map & struct:
	// map[string]person: lưu danh sách học sinh
	// 			{
	// 	"Duong" : person{
	// 		height: 178,
	// 		weight: 70,
	// 		fullName: "Dinh Hai Duong"
	// 	},
	// 	"Hieu": person{
	// 		...
	// 	}
	// }
	// fmt.Println(map["Duong"])

}
