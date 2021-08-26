package main

import "fmt"

type person struct {
	first      string
	last       string
	favflavors []string
}

func main() {
	HD := person{
		first: "Duong",
		last:  "Dinh",
		favflavors: []string{
			"caramen",
			"chocolate",
			"vani",
		},
	}
	TH := person{
		first: "Hieu",
		last:  "Nguyen",
		favflavors: []string{
			"maybe chocolate",
		},
	}
	fmt.Println(HD, TH)
	for i, v := range HD.favflavors {
		fmt.Println(i, v)
	}
	for i, v := range TH.favflavors {
		fmt.Println(i, v)
	}

	fmt.Println(TH)
	m := map[string]person{
		HD.last: HD,
		TH.last: TH,
		// "James": person {
		// 	first: "James",
		// 	last: "Bond",
		// 	favflavors: []string{
		// 		"Martini",
		// 		"Vani",
		// 	},
		// },
	}
	for _, v := range m {
		fmt.Println(v)
		for i, v := range v.favflavors {
			fmt.Println(i, v)
		}
	}
	// v= HD, HD.favflavors
	// v= TH, TH.favflavors
	// map không sắp xếp
}
