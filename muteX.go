package main

import (
	"fmt"
)

var value = 20

func muteX() {

	for i := 0; i < 5; i++ {
		wg.Add(2)
		mtx.Lock()
		go showValue()

		mtx.Lock()
		go increment()

		wg.Wait()
	}

}

func showValue() {
	defer mtx.Unlock()

	fmt.Println(value)

	wg.Done()

}
func increment() {
	defer mtx.Unlock()

	value++

	wg.Done()

}
