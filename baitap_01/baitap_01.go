package main

import "fmt"

func main() {
	M := 36
	N := 18
	for M != N {
		if M == N {
			break
		} else if M > N {
			M = M - N
		} else {
			N = N - M
		}
	}
	fmt.Println("UCLN cua M, N la", M)
}
