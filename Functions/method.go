// OOP

package main

import "fmt"

type human struct {
	name string
	age int
}

func (t human) eat()  {  // receiver có thể truy cập đến mọi trường (fields) và hành động (method) của type
	fmt.Println(t.name, "đang ăn")
}

func (t human) speakAge()  {
	t.eat()
	fmt.Println("Em đã", t.age)
}

// JS: func cũng có thể là 1 object
// func nào là object: viết hoa chữ đầu
// func nào là method: không viết hoa chữ cái đầu

func main()  {
	Duong := human{
		name: "Dương",
		age: 18,
	}
	fmt.Println(Duong)
	Duong.eat()
	Duong.speakAge()
}