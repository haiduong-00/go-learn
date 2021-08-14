package main

import "fmt"

func main() {
	S := 0
	for i := 9; i < 21; i++ {
		for 9 <= i < 11 {
			fmt.Println("S", S+i*30*120)
		}
		for 11 <= i < 13 ;i++ {
			fmt.Println("S", S+i*120*120)
		}
		for 13 <= i < 18 {
			fmt.Println("S", S+i*40*120)
		}
		for 18 <= i < 21 {
			fmt.Println("S", S+i*150*120)
		}
	} 
	fmt.Println("S")

}
