package main
import("fmt")
type company struct { 
	name string
	productField string
	stock float64
	tier string
}

func main() {
	Apple :=company {
		name :"Apple ",
		productField :"Technology" ,
		stock :148.19 ,
		tier :"Premier" ,
	}
	CocaCola := company{
		name :"Coca-cola",
		productField :"Berverage",
		stock: 56.64,
		tier :"B",
	}
	VinGroup:=company{
		name: "VinGroup",
		productField:"Apartment",
		stock:4.21,
		tier:"C",
	}
	fmt.Println(Apple)
	fmt.Println(CocaCola)
	fmt.Println(VinGroup)
}