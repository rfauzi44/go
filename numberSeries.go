package main

type numberSeries struct {
	Limit *int
}

func (n numberSeries) Even() []int {
	var evens []int
	for i := 1; i <= *n.Limit; i++ {
		if i%2 == 0 {
			evens = append(evens, i)
		}

	}
	return evens
}

func (n numberSeries) Odd() []int {
	var odds []int
	for i := 1; i <= *n.Limit; i++ {
		if i%2 != 0 {
			odds = append(odds, i)
		}

	}
	return odds
}

func (n numberSeries) Primes() []int {
	var primes []int
	for i := 2; i <= *n.Limit; i++ {
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

func (n numberSeries) Fib() []int {
	var fibSeq []int
	a, b := 0, 1
	for i := 0; i < *n.Limit; i++ {
		if a >= *n.Limit {
			break
		}
		fibSeq = append(fibSeq, a)
		a, b = b, a+b
	}
	return fibSeq
}
