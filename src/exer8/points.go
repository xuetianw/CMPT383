package exer8

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
