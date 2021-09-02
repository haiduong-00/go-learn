package main

import (
	"fmt"
)

func main() {
	i := 12
	switch {
	case i > 13:
		fmt.Println("có giá trị thảo mãn")
	case i < 10:
		fmt.Println("có giá trị ")
	}
}
