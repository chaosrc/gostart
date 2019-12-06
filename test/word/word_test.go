package word

import (
	"testing"
	
)

func TestPalindrome(t *testing.T) {
	if !IsPalindome("detartrated") {
		t.Error(`IsPalindome("detartrated") = false`)
	}

	if !IsPalindome("kayak") {
		t.Error(`IsPalindome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindome("palindrome") {
		t.Error(`IsPalindome("palindrome") = true`)
	}
}

func TestFranchPalindrome(t *testing.T) {
	if !IsPalindome("été") {
		t.Error(`IsPalindome("été") = false`)
	}
}
func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindome(input) {
		t.Errorf(`IsPalindome(%q) = false`, input)
	}
}

func BenchmarkPalindrome(b *testing.B) {
	for i := 0; i < b.N ; i++ {
		IsPalindome("A man, a plan, a cannal: panama")
	}
}