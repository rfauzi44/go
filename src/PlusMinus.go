package src

import "fmt"

func PlusMinus(arr []int32) {
	var pos int
	var neg int
	var zero int
	for i := 0; i < (len(arr)); i++ {
		if arr[i] > 0 {
			pos = pos + 1
		} else if arr[i] < 0 {
			neg = neg + 1
		} else {
			zero = zero + 1
		}
	}
	fmt.Printf("%.6f\n", float64(pos)/float64(len(arr)))
	fmt.Printf("%.6f\n", float64(neg)/float64(len(arr)))
	fmt.Printf("%.6f\n", float64(zero)/float64(len(arr)))

}
