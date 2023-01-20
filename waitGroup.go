package main

import (
	"fmt"
)

func waitGP() {
	var ch = make(chan int, 2)

	ch <- 10
	i := <-ch
	fmt.Println(i)
	ch <- 11
	j := <-ch
	fmt.Println(j)

	// wg.Add(2)

	// wg.Wait()

}

// func satu(ch chan int) {
// 	ch <- 1
// 	wg.Done()

// }
// func dua(ch chan int) {
// 	ch <- 2
// 	wg.Done()
// }
