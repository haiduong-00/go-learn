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
		panic(err)
	}
}

func main() {
	fo, err := os.Create("out.json")
	check(err)
	defer fo.Close()
	// Marshal (json) + write
	// Encode
	json.NewEncoder(fo).Encode(42)
	byteIn := make([]byte, 2)
	fo.Read(byteIn)
	fmt.Println(byteIn)
	var numIn int
	err = json.Unmarshal(byteIn, &numIn)
	check(err)
	fmt.Println(numIn)
}

// Open: read, execute
// Create: tao moi 1 file trong: write read execute
