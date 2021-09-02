package main

import (
	"fmt"
)

func main() {
	favSport := "badminton"
	switch favSport {
	case "football":
		fmt.Println("i played it for 7 year")
	case "badminton":
		fmt.Println("i played it for 5 year")
	}
}
