package exercise9

import (
	"math"
	"math/rand"
	"time"
)

type Sum struct {
	part1 float64
	part2 float64
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
		//fmt.Println(temp,ch)
		go calculation(temp, ch)
	}
	//fmt.Println(arr)

	var sum1 float64 = 0
	var squrSum float64 = 0

	for i := 0; i < chunks; i++ {
		p := <-ch
		sum1 += p.part1
		squrSum += p.part2
	}
	//close(ch)

	mean = (sum1 / float64(len(arr)))

	part1 := (squrSum) / float64(len(arr))
	part2 := sum1 / float64(len(arr)) * sum1 / float64(len(arr))

	stddev = math.Sqrt(float64(part1 - part2))

	return

}

func calculation(arr []int, sums chan Sum) {
	//fmt.Println(sums)
	var sum, sumsquare float64
	for i := 0; i < len(arr); i++ {
		sum += float64(arr[i])
		sumsquare += float64((arr[i]) * (arr[i]))
	}
	sums <- Sum{sum, sumsquare}
}

//func calculatesum(a Partialsum, arr []int, ch chan Partialsum) {
//	var sum,sumsquare float64
//	for i :=0; i<len(arr); i++ {
//		sum += float64(arr[i])
//		sumsquare += float64(arr[i])*float64(arr[i])
//	}
//	ch <- Partialsum{sum,sumsquare}
//}
