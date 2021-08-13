package main

import "fmt"

func main() {
	M := 15
	N := 5
	for {
		if M == N {
			fmt.Println("UCLN", M)
			break
		} else if M > N {
			M = M - N
		} else {
			N = N - M
		}
	}
}
