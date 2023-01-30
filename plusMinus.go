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
