package main

import "fmt"

func main() {
	for i := 9; i <= 21; i = i + 2 {
		for j := 9; j <= 11; j++ {
			fmt.Println(j * 30 * 120)
		}
	}
}
