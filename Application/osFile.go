// os package: package mà giao tiếp trục tiếp với hệ thống
// os thường nhập vào để xử lý file từ bên ngoài, xuất file

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	fo, err := os.Open("hello.json")
	defer fo.Close()
	check(err)
	var numIn int
	err = json.NewDecoder(fo).Decode(&numIn)
	check(err)
	fmt.Printf("%v %T\n", numIn, numIn)
	fmt.Println(fo.Name())
}
