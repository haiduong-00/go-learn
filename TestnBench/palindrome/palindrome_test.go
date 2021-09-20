package palindrome

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		desc   string
		N      int
		Output string
	}{
		{
			desc:   "123",
			N:      123,
			Output: "No",
		}, {
			desc:   "555",
			N:      555,
			Output: "Yes",
		},
		{
			desc:   "404",
			N:      404,
			Output: "Yes",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			re := IsPalindrome(tC.N)
			if re != tC.Output {
				t.Error("expected ", tC.Output, ", got ", re)
			}
		})
	}
}

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome(555))
	
	//Output: Yes
}

func Example() {
	fmt.Println(IsPalindrome(555))

	//Output: Yes
}