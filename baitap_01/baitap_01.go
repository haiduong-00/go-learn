package main

import "fmt"

func main() {
	for i := 10; i <= 21; i = i + 2 {
		for i := 10; i <= 11; i++ {
			fmt.Println(i * 30 * 120)
		}
		for i := 12; i <= 13; i++ {
			fmt.Println(i * 120 * 120)
		}
		for i := 14; i <= 18; i++ {
			fmt.Println(i * 40 * 120)
		}
		for i := 19; i <= 21; i++ {
			fmt.Println(i * 150 * 120)
		}
	}
}
