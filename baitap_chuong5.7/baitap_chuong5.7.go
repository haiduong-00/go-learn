package main

import "fmt"

func main() {
	x := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}
	for i, v := range x {
		fmt.Println(i, v)
		for i, v := range v {   // Block-level scope
			fmt.Println(i, v)	// Tìm i,v trong block trước, rồi tìm đến bên ngoài (nếu cần thiết)
		}
	}
}
