package main

import (
	"fmt"
)

// Non-anonymous struct

type person struct {
	height   int
	weight   int
	fullName string
}

type Duong struct {
	person              // Embedded type
	footballSkillRanked int
	footballPassion     bool
	fullName            string
}

type dog struct {
	isDog bool
}

// type <tên kiểu> struct {
// 		<tên đặc điểm> <kiểu dữ liệu>
// 		<tên đặc điểm> <kiểu dữ liệu>
// }

func (t *Duong) fastSprint() { // Method & Pointer
	fmt.Println("Dương đang dốc bóc rất nhanh đến khu vực đội bạn")
}

func (t *person) eat() {
	fmt.Printf("%v đang ăn\n", t.fullName)
}

func (t *dog) eat() {
	fmt.Println("Jake đang ăn")
}

func main() {
	Duong := Duong{ // Object: Duong
		person: person{ // Embedded struct
			height:   178,
			weight:   70,
			fullName: "Đinh Hải Dương",
		},
		footballSkillRanked: 8,
		footballPassion:     true,
		fullName:            "ABC",
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

	Jake := struct {
		dog  // Embedded struct
		age  int
		name string
		// isDog bool
	}{
		dog: dog{
			isDog: true,
		},
		age:  15,
		name: "Jake",
		// isDog: false,
	}
	// Anonymous struct so với OOP: (hỗ trợ -> có)
	// field, data: hỗ trợ đầy đủ
	// method, behavior: không hỗ trợ
	// Encapsulation: không hỗ trợ (do không có method, behavior)
	// Abstraction: chưa rõ
	// Inheritance: hỗ trợ đầy đủ (cả data và method)
	// Polymorphism: hỗ trợ 1 phần (chỉ data)

	// Polymorphism chỉ có ở field (data)
	Jake.eat()
	fmt.Println(Jake)
}
