package main

type smallStruct struct {
	a,b int
	c,d float64
}

//go:noinline
func smallAllocation() *smallStruct {
	return &smallStruct{}
}	
