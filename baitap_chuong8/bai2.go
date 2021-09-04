package main

import "fmt"

type person struct {
	first string
	age   int
}

func changeMe(p *person) {
	p.first = "duong"

}

func main() {
	Duong := person{
		first: "Dinh",
		age:   18,
	}
	fmt.Println(Duong)
	changeMe(&Duong)
	fmt.Println(Duong)
}
