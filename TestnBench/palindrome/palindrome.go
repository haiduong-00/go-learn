// Package palindrome để check dãy palind
package palindrome

// IsPalindrome check xem số đó có phải palindrome không
func IsPalindrome(n int) string {
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