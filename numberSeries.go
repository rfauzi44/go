package main

type numberSeries struct {
	limit *int
}

func (n numberSeries) even() []int {
	var evens []int
	for i := 1; i <= *n.limit; i++ {
		if i%2 == 0 {
			evens = append(evens, i)
		}

	}
	return evens
}

func (n numberSeries) odd() []int {
	var odds []int
	for i := 1; i <= *n.limit; i++ {
		if i%2 != 0 {
			odds = append(odds, i)
		}

	}
	return odds
}

func (n numberSeries) primes() []int {
	var primes []int
	for i := 2; i <= *n.limit; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}

func (n numberSeries) fib() []int {
	var fibSeq []int
	a, b := 0, 1
	for i := 0; i < *n.limit; i++ {
		if a >= *n.limit {
			break
		}
		fibSeq = append(fibSeq, a)
		a, b = b, a+b
	}
	return fibSeq
}
