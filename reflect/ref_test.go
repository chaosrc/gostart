package main

import (
	"testing"
	"os"
	"io"
)

func TestTypeof(t *testing.T) {
	var w io.Writer = os.Stdout
	t.Log(typeof(w))
}
func TestValueof(t *testing.T) {
	t.Log(valueof(1))
}
