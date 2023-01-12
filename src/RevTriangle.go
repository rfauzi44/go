package src

import "fmt"

func RevTriangle(row int) {
	firstRow := row + (row - 1)
	for i := 0; i <= row; i++ {
		for j := 0; j <= firstRow; j++ {
			if j <= firstRow-i && j >= i+1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
