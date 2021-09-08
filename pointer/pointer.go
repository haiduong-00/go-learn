package main

import (
	"fmt"
)
func changeMe(n *float64) {
	*n = 6
}

type person struct {
	name string
	age  int
}

type dog struct {
	age int
}

func (t person) speak() {
	fmt.Println("Hello, I'm", t.name, "and I'm", t.age)
}

func (t *person) aging() { // yêu cầu 1 pointer cho t
	t.age++
}

func (t *dog) aging() {
	t.age += 2
}


func (t person) run() {
	fmt.Println("run")
}

func (t dog) run() {
	fmt.Println("run now dog")
}

type animals interface { // Method sets: chứa person dog
	aging() // Method sets: của type *person và *dog
	run()
}

func getAge(t animals) {
	switch x := t.(type) {
	case *dog:
		fmt.Println(x.age)
	case *person:
		fmt.Println(x.age)
	}
	t.(*person).speak()
}

func run(t animals) {
	t.run()
}

// int string bool : nó không có 1 method sets nào
func main() {
	// var a float64
	// // a := 2
	// fmt.Println("before:", a)
	// b := &a
	// fmt.Println(*b)
	// *b = 5
	// fmt.Println("after:", a)

	// // b := a

	// fmt.Printf("%T %T %T\n", a, &a, *&a)

	// changeMe(&a)
	// fmt.Println("after:", a)

	// *T pointer của type T
	// *: kí hiệu đây là pointer
	// T: type của giá trị ở cái địa chỉ đấy

	// Nhập vào dữ liệu, data: nếu dữ liệu đấy lớn quá
	// mà trong Go, trong ngôn ngữ khác: muốn gán data sang identifer khác
	// bình thường Go lại copy nguyên cái data sang, và nó tốn bộ nhớ

	// Khi mà muốn thay đổi trực tiếp giá trị trong a

	// Method sets
	//1 struct, fields, method:
	// Interface: tập hợp các struct có method tên giống nhau

	Duong := person{
		name: "Duong",
		age:  17,
	}
	Jake := dog{
		age: 5,
	}
	Duong.aging()
	Jake.aging()
	getAge(&Duong)
	getAge(&Jake)
	run(&Duong)
}
