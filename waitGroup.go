package main

import (
	"fmt"
)

func waitGP() {
	var ch = make(chan int, 2)
	ch <- 10
	i := <-ch
	fmt.Println(i)

	// wg.Add(2)

	// wg.Wait()

}

func rizal(ch chan string) {
	ch <- "rizal"
	wg.Done()

}
func aldi(ch chan string) {
	ch <- "aldi"
	wg.Done()
}
