package main

import (
	"fmt"
)

//Exploring type
//Go is all about the type
var y int
var s = "\u65e5\u672c\u8a9e"
var z = `\u65e5\u672c\u8a9e`

//Go is a static programming language: C++ Pascal
//JavaScript is a dynamic programming language: python
// y ="Hello" print(Type)==String         y=42 print(Type)==int
//Go is precise

//Go: 2 kind of types
//1. Primitive data types: general, basic type, or built-in type
//2. Composite data types: compose other data

func main() {
	fmt.Println("Hello, playground")
	//In the flow of this doc, y is int
	y = 62
	fmt.Println(y)
	fmt.Printf("%v\n%T\n", s, s)
	fmt.Printf("%v\n%T", z, z) //formatting string
}
