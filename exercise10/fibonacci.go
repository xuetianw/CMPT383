package exercise10

func Fib(n uint, cutoff uint) uint {
	if n > cutoff {
		var ch = make(chan uint, 1)
		FibCocurrent(n, ch)
		return <-ch
	} else {
		return FibSequential(n)
	}
}

func FibSequential(u uint) uint {
	if u == 0 || u == 1 {
		return 1
	} else {
		return FibSequential(u-1) + FibSequential(u-2)
	}
}

func FibCocurrent(u uint, ch chan uint) {
	if u == 0 || u == 1 {
		ch <- 1
		return
	}
	leftChan := make(chan uint, 1)
	rightChan := make(chan uint, 1)
	go FibCocurrent(u-1, leftChan)
	go FibCocurrent(u-2, rightChan)
	left := <-leftChan
	right := <-rightChan
	ch <- left + right
	return

}

func Fibonacci(n uint) uint {
	return Fib(n, 30)
}

//
//func main() {
//	fmt.Print(Fibonacci(30))
//}
