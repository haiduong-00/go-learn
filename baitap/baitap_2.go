package main

import "fmt"

func main() {
	//Tính và in ra kết quả phép tính: S = (a+0) + (a+1) + (a+2) + ... + (a+n)
	// với 1/(a+n) > 0.0001
	a := 1.
	S := 0. //int:0 float64: 0.0 string: "" với những thứ khác: nil
	//single condition of for statements
	n := 0.
	for {
		S += (a + n)
		n++
		if 1/(a+n) < 0.0001 {
			break
		}
	}
	fmt.Println(S)
}
