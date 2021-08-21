package main

import (
	"fmt"
)

func main() {
	x := []int{2012, 2013, 2014, 2015, 2016, 2017, 2018, 2019, 2020, 2021}
	for _, v:= range x {
		switch {
		case v >= 2019:
			fmt.Println(v, "covid")
		}
	}
}
