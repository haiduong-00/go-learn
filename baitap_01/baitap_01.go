package main

import "fmt"

func main() {
	for i := 10; i <= 21; i = i + 2 {
		for i := 1; i <= 2; i++ {
			fmt.Println(i * 30 * 120)
		}
		for i := 1; i <= 2; i++ {
			fmt.Println(i * 120 * 120)
		}
		for i := 1; i <= 5; i++ {
			fmt.Println(i * 40 * 120)
		}
		for i := 1; i <= 3; i++ {
			fmt.Println(i * 150 * 120)
		}
	}
}
