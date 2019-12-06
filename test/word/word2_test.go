package word_test

import (
	"fmt"
	"testing"
	"example.com/hello/share/go-val"
)
func TestPP(t *testing.T) {
	fmt.Println(("ssccess"))
	val.Add()
	fmt.Println(val.Get())
}