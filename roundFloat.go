package main

import (
	"fmt"
	"math"
)

func roundFloat(val float64) {
	ratio := math.Pow(10, float64(1))
	result := math.Round(val*ratio) / ratio
	fmt.Printf("%.2f\n", result)
}
