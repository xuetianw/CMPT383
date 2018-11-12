package main

import (
	"fmt"
	"math"
)

func Hailstone(x uint) uint {
	if x%2 == 0 {
		return x / 2
	} else {
		return 3*x + 1
	}
}

func HailstoneSequenceAppend(a uint) []uint {
	list := []uint{}
	list = append(list, a)
	for list[len(list)-1] != 1 {
		list = append(list, Hailstone(list[len(list)-1]))
	}
	return list

}

func countLenth(x uint) uint {
	if x == 1 {
		return 1
	} else {
		return 1 + countLenth(Hailstone(x))
	}
}

func HailstoneSequenceAllocate(a uint) []uint {

	list := make([]uint, countLenth(a))
	list[0] = a
	for i := 1; i < len(list); i++ {
		temp := list[i-1]
		list[i] = Hailstone(temp)
	}
	return list
}

type Point struct {
	x float64
	y float64
}

func (p Point) String() string {
	return fmt.Sprintf("%v %v", p.x, p.y)
}

func NewPoint(x float64, y float64) Point {
	return Point{x, y}

}

func (p Point) Norm() float64 {
	return math.Sqrt((p.x)*(p.x) + (p.y)*(p.y))
}

func main() {

	fmt.Println(Hailstone(17))
	fmt.Println(Hailstone(18))

	fmt.Println(HailstoneSequenceAppend(18))
	a := []string{"a", "b"}
	fmt.Println(len(a))
	fmt.Println(a)
	a = append(a, "c")

	fmt.Println(len(a))
	fmt.Println(a)
	fmt.Println(a[0])

	fmt.Println(HailstoneSequenceAllocate(18))

	fmt.Println(countLenth(5))
	pt := NewPoint(3, 4.5)
	fmt.Println(pt) // should print (3, 4.5)
	pt = NewPoint(3, 4)
	fmt.Println(pt.Norm() == 5.0)

}
