package structI

import (
	"fmt"
)

func Start() {
	a := &Position{1, 2}
	// fmt.Println(a, zoom(a, 10))
	move(a, 3, 3)
	fmt.Println(a)

	b := Position{1, 2}
	c := Position{1, 3}
	z := Position{1, 2}

	fmt.Println(b == c) // false
	fmt.Println(b == z) // true

	fmt.Println("----------")
	cl := Circle{Position: Position{1, 1}, r: 2}
	// cl.p.x = 10
	fmt.Printf("%#v\n", cl)
	fmt.Println(cl.X)
	
	
	

}

// T t
type T struct {
	a int
	B string
}

// Position p
type Position struct {
	X, y int
}

func zoom(p Position, f int) Position {
	return Position{p.X * f, p.y * f}
}

func move(p *Position, x, y int) {
	p.X += x
	p.y += y
}

// Circle c
type Circle struct {
	Position
	r int
}
