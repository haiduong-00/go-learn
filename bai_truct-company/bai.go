package main
import("fmt")
type company truct { 
	name string
	productField string
	stock float64
	tier string
}
func main() {
	company :=Apple {
		name :"Apple ",
		productField :"Technology" ,
		stock :148.19 ,
		tier :"Premier" ,
	}
	company:=Coca-cola{
		name :"Coca-cola",
		productField :"Berverage",
		stock: 56.64,
		tier :"B",
	}
	company:=VinGroup{
		name: "VinGroup",
		productField:"Apartment",
		stock:4.21,
		tier:"C",
	}
	fmt.Println(Apple)
	fmt.Println(Coca-cola)
	fmt.Println(VinGroup)
}