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
	printfavflavors(HD)
	printfavflavors(TH)

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
		printfavflavors(v)
	}
	// v= HD, HD.favflavors
	// v= TH, TH.favflavors
	// map không sắp xếp
}

func printfavflavors(Obj person) {
	for i, v := range Obj.favflavors {
		fmt.Println(i, v)
	}
}
