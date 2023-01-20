package main

import (
	"fmt"
)

func channel() {

	wg.Add(2)
	var ch = make(chan int, 2) // bufer chanel

	go rOne(ch)
	go rTwo(ch)

	wg.Wait()

}

/*
Main \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
Routen 1 \\\\\\\\\\\\\\\\\
Routen 2 \\\\\\\\\\\\\\\\\
*/

// get data
func rOne(ch chan int) {
	// for data := range ch {
	// 	fmt.Println("datanya = ", data)
	// }

	i := <-ch
	fmt.Println(i)

	wg.Done()
}

// routine send data
func rTwo(ch chan int) {
	defer func() {
		close(ch)
		wg.Done()
	}()

	ch <- 20 // write channel
	fmt.Println("kirim")
	ch <- 50
}
