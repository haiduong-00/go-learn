package main

import (
	"fmt"
)

func main() {
	x := [10]int{2012, 2013, 2014, 2015, 2016, 2017, 2018, 2019, 2020, 2021}
	for i, v := range x {
		fmt.Println(i, v)
	}
}
