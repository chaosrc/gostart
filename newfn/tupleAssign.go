package main

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a %b
	}
	return a
}

func fib(n int) int {
	x, y := 0, 1
	
	for i := 2; i <= n ; i++ {
		x, y = y, x + y
	}
	return y
}