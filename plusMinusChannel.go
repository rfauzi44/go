package main

import "fmt"

func plusMinusChannel(arr []int, ch chan int) {
	defer func() {
		close(ch)
		wg.Done()
	}()
	var positive int
	var negative int
	var zero int
	for i := 0; i < (len(arr)); i++ {
		if arr[i] > 0 {
			positive = positive + 1
		} else if arr[i] < 0 {
			negative = negative + 1
		} else {
			zero = zero + 1
		}
	}
	ch <- positive
	ch <- negative
	ch <- zero

	i := <-ch
	j := <-ch
	k := <-ch
	fmt.Println(i, j, k)

}
