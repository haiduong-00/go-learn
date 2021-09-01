package main

import "fmt"

type person struct {
	first string
	last  string
	age   string
}

func (r person) speak() {
	fmt.Println(r.first ,r.last,r.age)
}
func main() {
	Duong:=person {
		first: "Duong",
		last: "Dinh",
		age : "almost18",
	}
	Duong.speak()
}