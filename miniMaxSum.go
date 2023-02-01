// Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers.

package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	var miniSum int64
	var maxSum int64
	for i := 0; i < len(arr)-1; i++ {

		miniSum = miniSum + int64(arr[i])

	}
	for i := 1; i < len(arr); i++ {

		maxSum = maxSum + int64(arr[i])

	}
	fmt.Printf("%v %v", miniSum, maxSum)
}
