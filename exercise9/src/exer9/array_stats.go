package exer9

import (
	"math/rand"
	"time"
)

func RandomArray(length int, maxInt int) []int {
	// TODO: create a new random generator with a decent seed; create an array with length values from 0 to values-1.

	rand.Seed(time.Now().UnixNano())
	randArr := []int{}
	for i := 0; i < length; i++ {
		randArr = append(randArr, rand.Intn(10)%maxInt)
	}
	return randArr

}

//func MeanStddev(arr []int, chunks int) (mean, stddev float64) {
//	if len(arr)%chunks != 0 {
//		panic("You promised that chunks would divide slice size!")
//	}
//	// TODO: calculate the mean and population standard deviation of the array, breaking the array into chunks segments
//	// and calculating on them in parallel.
//
//}
