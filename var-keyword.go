package main

import (
	"fmt"
)

var y int //Hey storage, I require a room for this identifier
//Zero value
//numeric = 0 float32: 0.00
//string = ""
//bool = false
//pointers arrays slices interfaces channels maps functions= nil
//var z [5]array
//z[1] =0
//array == nil

//var <identifier-name> <identifer-type> = <value>

func main() {
	//declare an identifier := : declare and ASSIGN
	// var : DECLARE (and assign)
	//scope of a variable
	//Flow of any programming language: from top to bottom, from left to right
	//Limit the scope of variables as much as possible + effective Go names = easy to read to understand to run
	//Use short declaration operator more
	boo()
	fmt.Println(y)
}

func boo() {
	x := 9
	var err error
	y, err = fmt.Println(x)
	fmt.Println(y)
	fmt.Println(err)
}

