package main

import (
	"fmt"
	"sort"
)

func findMedian(arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	n := len(arr)

	var median int32
	if n%2 == 0 {
		median = int32(arr[n/2-1]+arr[n/2]) / 2.0
	} else {
		median = int32(arr[n/2])
	}

	return median

}

func main() {
	arr := []int32{5, 3, 1, 2, 4}
	output := findMedian(arr)
	fmt.Println(output)

}
