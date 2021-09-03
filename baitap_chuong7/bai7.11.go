package main

import "fmt"

func recursion(x int) int {
	if x == 1 {
		return 1
	} else {
		return x * recursion(x-1)
	}
}
func main() {
	x := 5
	fmt.Println(recursion(x))
}