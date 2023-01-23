package main

import (
	"sync"
)

var wg = sync.WaitGroup{}
var mtx = sync.Mutex{}

func main() {
	// // Nomor 0
	// a := []int{7, 10, 2, 34, 33, -12, -8, 10}
	// chn := make(chan int)
	// go sum(a[:len(a)/2], chn)
	// go sum(a[len(a)/2:], chn)
	// x := <-chn
	// y := <-chn
	// fmt.Println(x + y)

	//Nomor 1
	// limit := 40
	// series := numberSeries{&limit}
	// wg.Add(4)
	// go series.even()
	// go series.odd()
	// go series.primes()
	// go series.fib()
	// wg.Wait()

	//Nomor 2
	// fibChan := make(chan []int)
	// wg.Add(2)

	// go fibbonaci(fibChan)
	// go evenOdd(fibChan)

	// wg.Wait()

	//Nomor 3
	// arr := []int{4, -1, 2, 0, 8, 9, 0, -8, 10}
	// ch := make(chan int, 3)
	// wg.Add(1)
	// go plusMinusChannel(arr, ch)
	// wg.Wait()

}
