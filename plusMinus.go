// Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero. Print the decimal value of each fraction on a new line with  places after the decimal.

package main

import "fmt"

func plusMinus(arr []int32) {
	var plus float32
	var minus float32
	var zero float32

	for _, v := range arr {
		if v > 0 {
			plus++
		} else if v < 0 {
			minus++
		} else {
			zero++

		}
	}

	fmt.Printf("%.6f\n", plus/float32(len(arr)))
	fmt.Printf("%.6f\n", minus/float32(len(arr)))
	fmt.Printf("%.6f\n", zero/float32(len(arr)))

}
