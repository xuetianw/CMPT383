package exer9

// TODO: Point (with everything from exercise 8 and) and methods that modify them in-place

import (
	"fmt"
	"math"
)

// TODO: The Point struct, NewPoint function, .String and .Norm methods
type Point struct {
	x float64
	y float64
}

func (p Point) String() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

func NewPoint(x float64, y float64) Point {
	return Point{x, y}

}
func (p Point) Norm() float64 {
	return math.Sqrt((p.x)*(p.x) + (p.y)*(p.y))
}
func (p *Point) Scale(i float64) {
	(*p).x *= i
	(*p).y *= i
}

func (p *Point) Rotate(a float64) {
	originX := (*p).x
	originY := (*p).y
	(*p).x = (*p).x*math.Cos(a) - (*p).y*math.Sin(a)
	(*p).y = originX*math.Sin(a) + originY*math.Cos(a)
}
