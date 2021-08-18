package main

import (
	"fmt"
)

func main() {
	x := 2003
	for {
		if x > 2021 {
			break
		}
		fmt.Println(x)
		x++
	}
}
