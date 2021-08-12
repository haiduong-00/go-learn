package main

import (
	"fmt"
)

var x int //data type int: integer

//type <type-name> <underlying-type>
type solieu1 string //data type

var y solieu1

//x : type: int8 - int64 { //Conversion
//	value: 19 (int8) - 19(int64) //Casting
//	???: type stored:
// 	storage:8bit - 64bit
//}

func main() {
	//var y solieuchofuncmain int
	//Go is all about the type
	//Dev can create their own type
	x = 19
	y = "19" //solieu1

	//Conversion, not Casting
	t := y == solieu1(fmt.Sprint(x)) //string
	//== : equality comparasion; = : assignment (same to C++)
	// === : (JavaScript) super equal to: JS is dynamic langiage, x=19 y="19"  x==y :true  x===y :false; Go is static language
	//Conversion only occur when 2 data is the same type kind
	
	fmt.Printf("%v %T\n%v %T\n%v", x, x, y, y, t)
}
