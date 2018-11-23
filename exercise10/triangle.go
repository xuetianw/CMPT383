package exercise10

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

type Triangle struct {
	A, B, C Point
}

func (t Triangle) String() string {
	return fmt.Sprintf("[%s %s %s]", t.A, t.B, t.C)
}

func (t *Triangle) Scale(i float64) {
	(*t).A.Scale(i)
	(*t).A.Scale(i)
	(*t).B.Scale(i)
}

func (t *Triangle) Rotate(a float64) {
	(*t).A.Rotate(a)
	(*t).B.Rotate(a)
	(*t).C.Rotate(a)
}

type Transformable interface {
	Scale(i float64)
	Rotate(i float64)
}

func TurnDouble(t Transformable, angle float64) {
	t.Scale(2)
	t.Rotate(angle)
}

//func main() {
//	pt := Point{3, 4}
//	TurnDouble(&pt, 3*math.Pi/2)
//	fmt.Println(pt)
//	tri := Triangle{Point{1, 2}, Point{-3, 4}, Point{5, -6}}
//	TurnDouble(&tri, math.Pi)
//	fmt.Println(tri)
//}
