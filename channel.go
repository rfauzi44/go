package main

import (
	"fmt"
)

func Fibonacci(ch chan []int) {

	result := []int{}
	n1 := 0
	n2 := 1
	n3 := n2

	for i := 1; i <= 20; i++ {
		result = append(result, n2)
		n3 = n1 + n2
		n1 = n2
		n2 = n3
		if n2 > 20 {
			break
		}
	}

	fmt.Println("Bilangan Fibonacci: ", result)

	ch <- result
	// close(ch)

	// Sender

	wg.Done()

}

func GanjilGenap(ch chan []int) {

	resultGanjil := []int{}
	resultGenap := []int{}

	fibo := <-ch

	for _, v := range fibo {
		if v%2 != 0 {
			resultGanjil = append(resultGanjil, v)
		}
		if v%2 == 0 {
			resultGenap = append(resultGenap, v)
		}

	}

	fmt.Print("Bilangan Ganjil: ", resultGanjil)
	fmt.Println()
	fmt.Print("Bilangan Genap: ", resultGenap)
	fmt.Println()

	// Receiver

	ch <- resultGanjil

	wg.Done()
}

func Sum(ch chan []int) {

	num := <-ch

	result := 0
	for _, v := range num {
		result += v

	}

	fmt.Println("Hasil sum: ", result)

	wg.Done()

}
