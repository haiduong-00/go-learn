package main

import "fmt"

func main() {
	M := 36
	N := 18
	for M != N {
		if M > N {
			M = M - N
		} else {
			N = N - M
		}
	}
	fmt.Printf("UCLN cua M, N la %v", M)
}
