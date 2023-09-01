package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и
// конструктором.

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func Distance(p1, p2 *Point) float64 {
	x2 := math.Pow(p2.x-p1.x, 2)
	y2 := math.Pow(p2.y-p1.y, 2)

	return math.Sqrt(x2 + y2)
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(3, 5)

	fmt.Println(Distance(p1, p2))
}
