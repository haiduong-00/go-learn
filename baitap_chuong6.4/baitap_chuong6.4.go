package main

import "fmt"

func main() {
	x := struct {
		first      string
		last       string
		favFlavors []string
	}{
		first: "Duong",
		last:  "Dinh",
		favFlavors: []string{
			"vani", "chocolate",
		},
	}
	fmt.Println(x)
}