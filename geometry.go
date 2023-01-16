package main

import "math"

type Calculate2d interface {
	Area() float64
	Circ() float64
}

type Calculate3d interface {
	Volume() float64
}

type Calculate interface {
	Calculate2d
	Calculate3d
}

type Circle struct {
	Radius *float64
}

type Square struct {
	Side *float64
}

func (c Circle) Area() float64 {
	return math.Phi * math.Pow(*c.Radius, 2)
}

func (c Circle) Circ() float64 {
	return (2 * math.Phi) * *c.Radius
}

func (c Circle) Volume() float64 {
	return ((4 / 3) * math.Phi) * math.Pow(float64(*c.Radius), 3)
}

func (s Square) Area() float64 {
	return math.Pow(*s.Side, 2)
}

func (s Square) Circ() float64 {
	return 4 * *s.Side
}

func (s Square) Volume() float64 {
	return math.Pow(*s.Side, 3)
}
