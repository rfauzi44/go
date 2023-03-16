package main

import "fmt"

// Given an array of integers and a positive integer , determine the number of  pairs where  and  +  is divisible by .
// divisible = 5
// i < j and ar[i] + ar[j] divisible by 5

func main() {
	var n int32 = 6
	var k int32 = 5
	var ar = []int32{1, 2, 3, 4, 5, 6}

	output := divisibleSumPairs(n, k, ar)
	fmt.Println(output)

}

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {

	pairs := 0

	for i := 0; i < int(n); i++ {

		for j := 1; j < int(n); j++ {

			if (ar[i]+ar[j])%k == 0 && i < j {
				pairs++

			}

		}

	}
	return int32(pairs)

}
