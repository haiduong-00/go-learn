package main

import (
	"fmt"
	"sort"
)

type student []struct {
	name  string
	score float64
}

func (t student) Len() int      { return len(t) }
func (t student) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t student) Less(i, j int) bool {
	if t[i].name < t[j].name {
		return true
	} else {
		return false
	}
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	dsHocsinh := student{
		{name: "Diep", score: 9.8},
		{name: "Hieu", score: 10},
		{name: "Duong", score: 8.6},
	}
	fmt.Println(dsHocsinh)
	sort.Sort(dsHocsinh)
	fmt.Println(dsHocsinh)
}
