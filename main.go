package main

import (
	"fmt"
)

func main() {
	roundFloat(4.37)
	var limit int = 50
	var series1 numberSeries
	series1 = numberSeries{Limit: &limit}
	fmt.Println(series1.Even())
	var radius float64 = 5
	var shape1 Calculate
	shape1 = Circle{Radius: &radius}
	fmt.Println(shape1.Area())
	// var shape2 *src.Calculate
	// shape2 = &shape1
	// fmt.Println(*shape2)
	var side float64 = 4
	var shape2 Calculate
	shape2 = Square{Side: &side}
	fmt.Println(shape2.Area())

}
