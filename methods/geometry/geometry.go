package geometry

import (
	"image/color"
	"math"
	"sync"
)

// Point 位置
type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	Path
	color.RGBA
}

// Distance 计算两点之间的距离
func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.Y-q.Y, p.X-q.X)
}

// Path 路径
type Path []Point

// Distance 计算 Path 的距离
func (p Path) Distance() float64 {
	var sum float64

	for i, point := range p {
		if i > 0 {
			sum += point.Distance(p[i-1])
		}
	}
	return sum
}

//Add 相加
func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }

//Substract 相减
func (p Point) Substract(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

// ScaleBy 放大
func (p *Point) ScaleBy(i float64) {
	p.X *= i
	p.Y *= i
}

// Line 线
type Line struct {
	Start Point
	End   Point
}

// Transition 移动
func (p Path) Transition(offset Point, add bool) Path {
	var op func(Point, Point) Point

	if add {
		op = Point.Add
	} else {
		op = Point.Substract
	}

	for i, point := range p {
		p[i] = op(point, offset)
	}
	return p
}

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func lookup(key string) string {
	cache.Lock()
	defer cache.Unlock()
	return cache.mapping[key]
}

func UseCache() string {
	return lookup("123")
}
