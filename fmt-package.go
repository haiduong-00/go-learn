package main

import (
	"fmt"
)

func main() {
	x := 32.2                    //raw == raw, interpreted -> raw
	y:= fmt.Sprintf("%g \\%T\"", x, x) //string format --> print this string to a identifier
	//%b: print the number in binary type (base 2)
	//%x: print the number in hexadecimal type (base 16) char in hex is lowercase
	//In float number:
	//%e: print float in 10^xx type( present exx) (scientific type)
	//%f: print float in decimal type (normal 0.0000 type)
	//%g: print automatically in %e or %f, automatically format
	//In string:
	//%q: print raw string, even if it's in interpreted string
	//Escape characters
	// \n: new line
	// \a: alert, bell
	// \b: backspace reverse 1 space
	// \f: form feed
	// \r:carriage return == \n
	// \t: tab
	// \": "
	// \\: \
	fmt.Print(y)
}
