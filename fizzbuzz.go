package main

import (
	"fmt"
)

/*
 * Complete the 'fizzBuzz' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func fizzBuzz(n int32) {
	for i := 1; i <= int(n); i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")

		} else if i%5 == 0 {
			fmt.Println("Buzz")

		} else {
			fmt.Println(i)
		}

	}

}

func main() {

	fizzBuzz(15)

}
