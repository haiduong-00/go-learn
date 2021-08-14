package main

import "fmt"

func main() {
<<<<<<< HEAD
	M := 9
	N := 18
	for M != N {
=======
	M := 15
	N := 5
	for {
>>>>>>> 0f15e8477303765884ed94bde54c3f9f92900f99
		if M == N {
			fmt.Println("UCLN", M)
			break
		} else if M > N {
			M = M - N
		} else {
			N = N - M
		}
	}
<<<<<<< HEAD

	fmt.Println("UCLN", M)

=======
>>>>>>> 0f15e8477303765884ed94bde54c3f9f92900f99
}
