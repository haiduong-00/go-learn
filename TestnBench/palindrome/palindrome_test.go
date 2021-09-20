package palindrome

import (
	"fmt"
	"testing"
)

func TestPalidrome_first(t *testing.T) {
	re := is_palindrome(555)
	if re != "Yes" {
		t.Error("expected Yes, got",re) // Log then Fail (continue)
	}
	fmt.Println("hello")
}

func TestPalidrome_second(t *testing.T) {
	re := is_palindrome(12321)
	if re != "Yes" {
		t.Error("expected Yes, got",re)
	}
}