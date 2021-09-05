package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type testIn []struct {
	N int `json:"N"`   //Change me
}

type testOut []struct {
	Output int `json:"Output"`  //Change me
}

func ReadIn() (testIn, testOut) {
	fi, e := os.Open("testIn.json")
	check(e)
	fo, e := os.Open("testOut.json")
	check(e)
	defer fi.Close()
	defer fo.Close()

	var testIn testIn
	var testOut testOut

	json.NewDecoder(fi).Decode(&testIn)

	json.NewDecoder(fo).Decode(&testOut)

	return testIn, testOut
}

func findNthTerm(N int) int {
	for _, v := range x {
		if N == 1 {
			fmt.Println(v)
		}else if N==2 {
			fmt.Println(v)
		}else if i==3 {
			fmt.Println(v)
		}else if N==4 {
			fmt.Println(v)
	}else if N==5 {
		fmt.Println(v)
	}else if N==6 {
		fmt.Println(v)
}
return N
}

func main(){}