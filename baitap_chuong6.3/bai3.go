package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}


func main() {
	a := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black and red",
		},
		fourWheel: true,
	}
	b := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "yellow and white",
		},
		luxury: true,
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a.color)
	fmt.Println(b.color)
}
