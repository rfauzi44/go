// Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers.

package main

import (
	"fmt"
	"sort"
)

func minMaxSum(arr []int32) {
	var minSum int64
	var maxSum int64
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	for i := 0; i < len(arr)-1; i++ {
		minSum = minSum + int64(arr[i])
	}
	for i := len(arr) - 1; i > 0; i-- {
		maxSum = maxSum + int64(arr[i])
	}
	fmt.Println(minSum, maxSum)

}
