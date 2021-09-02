package main

import "fmt"

func is_palindrome(n int) string {
	x := []int{}
	for n >= 10 {
		x = append(x, n%10)
		n = n / 10
	}
	x = append(x, n)
	for i := len(x) / 2; i >= 0; i-- {
		if x[i] == x[len(x)-1-i] {
		} else {
			return "No"
		}
	}
	return "Yes"
}
func main() {
	n := 12321
	fmt.Println(is_palindrome(n))
}
