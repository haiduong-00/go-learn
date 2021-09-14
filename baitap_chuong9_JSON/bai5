package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}
type Age []user

	func(t Age)Len()int{return len(t)}
	func(t Age)Swap(i,j int){t[i],t[j]=t[j],t[i]}
	func(t Age)Less(i,j int)bool{return  t[i].Age<t[j].Age }
			
		
	
	type Last []user
	
	func (a Last) Len() int           { return len(a) }
	func (a Last) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
	func (a Last) Less(i, j int) bool { return a[i].Last < a[j].Last }
	 	
func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	for _,a:=range users{
		fmt.Println(a.First,a.Last,a.Age)
		sort.Strings(a.Sayings)
		for _,v:=range a.Sayings {
			fmt.Println("\t",v)
		}
	}
	// for _,u:=range users{
	// 	fmt.Println(u.First,u.Last,u.Age)
	// 	sort.Strings(u.Sayings)
	// 	for _,v:=range u.Sayings{
	// 		fmt.Println(v)
	// 	}
	// }
	fmt.Println("----------------------------------------------------------------------")
	sort.Sort(Age(users))
	for _,a:=range users{
		fmt.Println(a.First,a.Last,a.Age)
		for _,v:=range a.Sayings {
			fmt.Println("\t",v)
		}
	}
	fmt.Println("----------------------------------------------------------------------")
	sort.Sort(Last(users))
	for _,a:=range users{
		fmt.Println(a.First,a.Last,a.Age)
		for _,v:=range a.Sayings {
			fmt.Println("\t",v)
		}
	}

	
	
	
	
	}

