package word

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindome("detartrated") = false`)
	}

	if !IsPalindrome("kayak") {
		t.Error(`IsPalindome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindome("palindrome") = true`)
	}
}

func TestFranchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindome("été") = false`)
	}
}
func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindome(%q) = false`, input)
	}
}

func BenchmarkPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	// Output: true
}
