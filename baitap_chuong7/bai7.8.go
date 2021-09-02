package main

import "fmt"

func main() {
	a := duong()
	fmt.Println(a())
}
func duong() func() int {
	return func() int {
		return 18
	}

}
