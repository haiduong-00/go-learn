package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`}}
	for k, v := range m {
		fmt.Println(k, v)
		for i, v := range v {
			fmt.Println(i, v)
		}
	}
	m["darcular"] = []string{"âmmama", "cmcmcm", "cmcdfe"}
	for k, v := range m {
		fmt.Println(k, v)
		for i, v := range v {
			fmt.Println(i, v)
		}
	}
	delete(m, "darcular")
	for k, v := range m {
		fmt.Println(k, v)
		for i, v := range v {
			fmt.Println(i, v)
		}
	}
}
