package palindrome

import (
	"testing"
)

func TestPalindrome(t *testing.T) {
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
			re := is_palindrome(tC.N)
			if re != tC.Output {
				t.Error("expected ", tC.Output, ", got ", re)
			}
		})
	}
}