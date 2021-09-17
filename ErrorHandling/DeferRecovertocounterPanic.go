package main

import "fmt"

// recover(): câu lệnh để:
// Cái này được sử dụng để giữ cho panic không gây panicking đến caller nào đó
// recover chỉ chạy khi đi kèm defer

func HowtoUseRecover() {
	foo()
	fmt.Println("Hello")
}

func foo() {
	err := fmt.Errorf("một lỗi nữa")
	defer func ()  {
		fmt.Println("Vẫn chạy")
		recover()
	}()
	check(err)
	fmt.Println("Có chạy không")
}