package exer9

import (
	"math"
	"math/rand"
	"time"
)

type Sum struct {
	part1 int
	part2 int
}

func RandomArray(length int, maxInt int) []int {
	// TODO: create a new random generator with a decent seed; create an array with length values from 0 to values-1.

	rand.Seed(time.Now().UnixNano())
	randArr := []int{}
	for i := 0; i < length; i++ {
		randArr = append(randArr, rand.Intn(10)%maxInt)
	}
	return randArr

}

func MeanStddev(arr []int, chunks int) (mean, stddev float64) {
	if len(arr)%chunks != 0 {
		panic("You promised that chunks would divide slice size!")
	}
	// TODO: calculate the mean and population standard deviation of the array, breaking the array into chunks segments
	// and calculating on them in parallel.

	var ch = make(chan Sum)

	size := len(arr) / chunks
	for i := 0; i < len(arr); i += size {
		temp := arr[i : i+size]
		go calculation(temp, ch)
	}
	close(ch)

	var sum1 int
	var squrSum int

	for item := range ch {
		sum1 += item.part1
		squrSum += item.part2
	}

	mean = float64(sum1 / len(arr))

	part1 := (squrSum) / len(arr)
	part2 := sum1 / len(arr)

	stddev = math.Sqrt(float64(part1 - part2))

	return

}

func calculation(ints []int, sums chan Sum) {
	var sum int
	for i := 0; i < len(ints); i++ {
		sum += ints[i]
	}
	var sumSqu int
	for i := 0; i < len(ints); i++ {
		sumSqu += (ints[i]) * (ints[i])
	}
	sums <- Sum{sum, sumSqu}
}
