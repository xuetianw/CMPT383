package main

import "fmt"

func Fib(n uint, cutoff uint) uint {
	var ch = make(chan uint, 1)
	FibCocurrent(n, ch, cutoff)
	return <-ch
}

func FibCocurrent(u uint, ch chan uint, cutoff uint) {
	if u == 0 || u == 1 {
		ch <- 1
		return
	}
	leftChan := make(chan uint, 1)
	rightChan := make(chan uint, 1)
	if (cutoff > 0) {
		go FibCocurrent(u - 1, leftChan, cutoff - 1)
		go FibCocurrent(u - 2, rightChan, cutoff -1)
	} else {
		FibCocurrent(u - 1, leftChan, cutoff - 1)
		FibCocurrent(u - 2, rightChan, cutoff - 1)
	}
	left := <-leftChan
	right := <-rightChan
	ch <- left + right
	return

}

func Fibonacci(n uint) uint {
	return Fib(n, 1000)
}

func main() {
	fmt.Print(Fibonacci(40))
}
