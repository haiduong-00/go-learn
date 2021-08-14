package main

import "fmt"

func main() {
<<<<<<< HEAD
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
=======
	for i := 10; i <= 21; i = i + 2 {
		for i := 1; i <= 2; i++ {
			fmt.Println(i * 30 * 120)
		}
		for i := 1; i <= 2; i++ {
			fmt.Println(i * 120 * 120)
		}
		for i := 1; i <= 5; i++ {
			fmt.Println(i * 40 * 120)
		}
		for i := 1; i <= 3; i++ {
			fmt.Println(i * 150 * 120)
>>>>>>> 0d3c324162bb59b22291c4fe010a6d03b5c67fa3
		}
	}
<<<<<<< HEAD

	fmt.Println("UCLN", M)

=======
>>>>>>> 0f15e8477303765884ed94bde54c3f9f92900f99
}
