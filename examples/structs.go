package main

import "fmt"

type rect struct {
	width  float64
	heigth float64
}

func (r rect) area() float64 {
	return r.width * r.heigth
}

type square struct {
	sides float64
}

func (s square) area() float64 {
	return s.sides * s.sides
}

func main() {
	s := rect{
		width:  54.0,
		heigth: 54.0,
	}

	cuadrado := square{
		sides: 3.0,
	}
	fmt.Println(s.area())
	fmt.Println(cuadrado.area())
}
