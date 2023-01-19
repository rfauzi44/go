package main

import (
	"fmt"
	"time"
)

var value = 20

func muteX() {

	for i := 0; i < 5; i++ {
		wg.Add(2)

		go showValue()
		time.Sleep(1 * time.Second)

		go increment()
		wg.Wait()
	}

}

func showValue() {

	mtx.Lock()
	fmt.Println(value)
	mtx.Unlock()
	wg.Done()
}
func increment() {

	mtx.Lock()
	value++
	mtx.Unlock()
	wg.Done()
}
