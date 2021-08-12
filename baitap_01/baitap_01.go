package main

import "fmt"

func main() {
	M := 1
	N := 5
	for M != N {
		if M == N {
			fmt.Println("UCLN", M)
			break
		} else if M > N {
			M = M - N
		} else {
			N = N - M
		}
	}

	fmt.Println("UCLN", M)

}
