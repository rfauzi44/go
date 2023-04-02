package main

import "fmt"

type numberSeries struct {
	limit *int
}

func (n numberSeries) even() {
	defer wg.Done()
	var evens []int
	for i := 1; i <= *n.limit; i++ {
		if i%2 == 0 {
			evens = append(evens, i)
		}

	}
	fmt.Println(evens)
}

func (n numberSeries) odd() {
	defer wg.Done()
	var odds []int
	for i := 1; i <= *n.limit; i++ {
		if i%2 != 0 {
			odds = append(odds, i)
		}

	}
	fmt.Println(odds)
}

func (n numberSeries) primes() {
	defer wg.Done()
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
	fmt.Println(primes)
}

func (n numberSeries) fib() {
	defer wg.Done()
	var fibSeq []int
	a, b := 0, 1
	for i := 0; i < *n.limit; i++ {
		if a >= *n.limit {
			break
		}
		fibSeq = append(fibSeq, a)
		a, b = b, a+b
	}
	fmt.Println(fibSeq)
}
