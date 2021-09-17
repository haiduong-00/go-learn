package main

import "fmt"

func main() {
	foo()
}

func foo() {
	err := fmt.Errorf("một lỗi nữa")
	defer fmt.Println("Vẫn chạy")
	check(err)
}