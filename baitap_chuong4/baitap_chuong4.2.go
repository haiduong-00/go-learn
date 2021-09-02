package main

import (
	"fmt"
)

func main() {
	for i := 0x41; i <= 0x5a; i++ {
		fmt.Printf("%d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
	x := []rune{
		0x0078,
		0x0069,
		0x006e,
		0x0020,
		0x0063,
		0x0068,
		0x00e0,
		0x006f,
		0x0020,
		0x1f1fb,
		0x1f1f3,
	} // String cũng chỉ là slice of byte
	fmt.Println(string(x))
}
