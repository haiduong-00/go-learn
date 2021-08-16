package main

import "fmt"

const favSport string = "football"
const callout int = 4

var (
	Jane int = 0
	Josh int = 1
	Jihn int = 2
)

func main() {
	//Bài yêu cầu kiểm tra 1 trong 3 callout, nếu callout nào đúng thì in ra callout đấy
	//if, 3 identifier, 2 constants,

	if callout == Jane {
		fmt.Println("Jane", favSport)
	} else if callout == Josh {
		fmt.Println("Josh", favSport)
	} else if callout == Jihn {
		fmt.Println("Jihn", favSport)
	} else {
		fmt.Println("ko có gì")
	}
}
