package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type student []struct {
	name  string
	score float64
}

func (t student) Len() int {
	return len(t)
}

func (t student) Less(i, j int) bool {
	if t[i].score < t[j].score {
		return true
	} else {
		return false
	}
}

func (t student) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	fi, err := os.Open("hello.json")
	check(err)
	defer fi.Close()
	var stringIn []string
	err = json.NewDecoder(fi).Decode(&stringIn)
	check(err)

	fmt.Println(stringIn)

	dsHocsinh := student{
		{name: "Duong", score: 9.5},
		{name: "Hieu", score: 9},
		{name: "James", score: 8},
		{"An", 8},
	}
	sort.Strings(stringIn)
	fmt.Println(stringIn)

	fmt.Println(dsHocsinh)
	sort.Sort(dsHocsinh)
	fmt.Println(dsHocsinh)
}
