package main

import (
	"fmt"
)

func main() {
	roundFloat(4.37)
	var limit int = 50
	var series1 numberSeries
	series1 = numberSeries{limit: &limit}
	fmt.Println(series1.even())
	var radius float64 = 5
	var shape1 calculate
	shape1 = circle{radius: &radius}
	fmt.Println(shape1.area())
	var side float64 = 4
	var shape2 calculate
	shape2 = square{side: &side}
	fmt.Println(shape2.area())

}
