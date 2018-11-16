package exer9

import (
	"math/rand"
)

// Partition the slice arr around a random pivot (in-place), and return the pivot location.
func partition(arr []float64) int {
	// Adapted from https://stackoverflow.com/a/15803401/6871666
	left := 0
	right := len(arr) - 1

	// Choose random pivot
	pivotIndex := rand.Intn(len(arr))

	// Stash pivot at the right of the slice
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Move elements smaller than the pivot to the left
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Place the pivot after the last-smaller element
	arr[left], arr[right] = arr[right], arr[left]
	return left
}

func InsertionSort(arr []float64) {
	// TODO: implement insertion sort

	for i := 1; i < len(arr); i++ {
		n := arr[i]
		j := i
		for j > 0 && arr[j-1] > n {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = n
	}
}

const insertionSortCutoff = 10000

func quicksort(arr []float64) {
	low := 0
	high := len(arr) - 1
	if high-low < 2 {
		p := partition(arr)
		quicksort(arr[low:p])
		quicksort(arr[p:high])
	}
}

func QuickSort(arr []float64) {
	// TODO: implement Quicksort:
	//   do nothing for length < 2
	//   do insertion sort for length < insertionSortCutoff
	//   do Quicksort otherwise.
	// TODO: decide on a good value for insertionSortCutoff

	if len(arr) < insertionSortCutoff {
		InsertionSort(arr)
	} else {
		quicksort(arr)
	}

}
