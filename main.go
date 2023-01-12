package main

import "fmt"

func main() {
	// var arr = [3]int{5, 6, 3}
	// fmt.Println(arr)
	// arr[0] = 6
	// var u int = arr[0]
	// u = 5
	// fmt.Println(u)

	// var a int = 5
	// var b *int = &a
	// fmt.Println(b)
	// c := *b
	// fmt.Println(c)

	type car struct {
		color string
		wheel int
	}

	var ferrari1 car
	ferrari1.color = "blue"
	ferrari1.wheel = 5

	var ferrari2 *car = &ferrari1

	fmt.Println(ferrari1)
	fmt.Println(ferrari2.wheel)
}
