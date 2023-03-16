package main

import (
	"fmt"
)

/*
 * Complete the 'matchingStrings' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY strings
 *  2. STRING_ARRAY queries
 */

func main() {

	var strings = []string{"aba", "baba", "aba", "xzxb"}
	var queries = []string{"aba", "xzxb", "ab"}
	output := matchingStrings(strings, queries)
	fmt.Println(output)

}

func matchingStrings(strings []string, queries []string) []int32 {

	var matchingArray = make([]int32, len(queries))

	for i := 0; i < len(queries); i++ {
		var matching int32 = 0

		for j := 0; j < len(strings); j++ {
			if queries[i] == strings[j] {
				matching++
			}

		}
		matchingArray[i] = matching

	}

	return matchingArray
	// Write your code here

}
