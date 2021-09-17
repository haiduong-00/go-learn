package main

import "fmt"

func main() {
	fmt.Println(foo(1))
}

func foo(i int) (y int) {
	defer func ()  {
		y++
	}()
	return i 
}