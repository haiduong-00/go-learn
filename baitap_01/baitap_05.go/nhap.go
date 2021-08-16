package main

import "fmt"

func main() {
	S := 0
	for i := 9; i < 21; i++ {
		if 9 <= i < 11 {
			fmt.Println(S + i*30*120)
		} else if 11 <= i < 13 {
			fmt.Println(S + i*120*120)
		} else if 13 <= i < 18 {
			fmt.Println(S + i*40*120)
		} else if 18 <= i < 21 {
			fmt.Println(S + i*150*120)
		}
	}
	fmt.Println("S")

}
