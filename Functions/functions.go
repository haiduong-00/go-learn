package main

import "fmt"

// func (r receiver) identifier(paramenters) (returns) { code }

func main()  {
	foo() // báo hiệu rằng đây là gọi function ra để CHẠY
	// bar := foo  // báo hiệu rằng gắn ĐOẠN CODE của foo cho bar
	// foo: (string:func ()  {
	// 	fmt.Println("Hello, World!")
	// } )
	hi := "Hello"
	Duong := "Duong"
	bar(Duong,hi,18)
	foo()
	tong,hieu := addAndsubtract(5,8)
	bar("là","Tổng",tong)
	bar("là","Hiệu",hieu)
}

func foo()  {
	fmt.Println("Hello, World!")
}

func bar(name, greetings string, num int)  { // greetings: Lời chào
	fmt.Println(greetings,name,num)
	// (identifier type, identifier type)
}

func addAndsubtract(num1, num2 int) (int, int) {  // Cộng 2 số num1, num2 lại và trả kết quả (value) là tổng
	return num1+num2,(num1-num2)
}

// Hồi xưa: đoạn code là các dòng trên 1 trang giấy và 1 chương trình là nguyên 1 tờ
// Code ngày càng nặng, nhiều chức năng nên người ta sáng tạo ra cách chia nhỏ code
// Function
// Package

// Có suy nghĩ build func từ đầu chia nhỏ code: sẽ biết được sai ở đâu