package main

import "fmt"

const favSport string = "football"
const callout int = 2

var (
	Jane int = 0
	Josh int = 1
	Jihn int = 2
)

func main() {
	if callout == Jane {
		fmt.Println("Jane", favSport)
	} else if callout == Josh {
		fmt.Println("Josh", favSport)
	} else {
		fmt.Println("Jihn", favSport)
	}
}
