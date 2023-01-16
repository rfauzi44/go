package main

import "fmt"

func searchArray() {
	var arr = []int{5, 6, 10, 8}
	for i := 0; i < len(arr); i++ {
		if arr[i] == 7 {
			fmt.Println("Thereis")
			break
		} else if i == len(arr)-1 {
			fmt.Println("None")
		}
	}

}
