package main

import "math"

type calculate2d interface {
	area() float64
	circ() float64
}

type calculate3d interface {
	volume() float64
}

type calculate interface {
	calculate2d
	calculate3d
}

type circle struct {
	radius *float64
}

type square struct {
	side *float64
}

func (c *circle) area() float64 {
	return math.Phi * math.Pow(*c.radius, 2)
}

func (c *circle) circ() float64 {
	return (2 * math.Phi) * *c.radius
}

func (c *circle) volume() float64 {
	return ((4 / 3) * math.Phi) * math.Pow(float64(*c.radius), 3)
}

func (s *square) area() float64 {
	return math.Pow(*s.side, 2)
}

func (s *square) circ() float64 {
	return 4 * *s.side
}

func (s *square) volume() float64 {
	return math.Pow(*s.side, 3)
}
