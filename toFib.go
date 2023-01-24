package main

import "fmt"

func fibbonaci(ch chan<- []int) {
	defer wg.Done()
	limit := 40
	series := numberSeriesReturn{&limit}
	seriesFib := series.fib()
	ch <- seriesFib
}

func odd(ch <-chan []int) {
	defer wg.Done()
	theFib := <-ch
	var result []int
	for _, v := range theFib {
		if v%2 != 0 {
			result = append(result, v)
		}
	}
	fmt.Println("ganjil", result)
}
