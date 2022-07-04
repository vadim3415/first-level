package main

import (
	"firstLevel/develop/dev24/pointer"
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и
// конструктором.

func main() {
	fmt.Println("расстояние между двумя точками: ", Distance())
}


func Distance() float64 {
	point1 := pointer.NewPoint(2, 1)
	point2 := pointer.NewPoint(1, -3)
	x1, y1 := point1.Get()
	x2, y2 := point2.Get()

	fmt.Println("x1, y1:", point1)
	fmt.Println("x2, y2:", point2)

	return math.Sqrt(math.Pow((x1-y1), 2) + math.Pow((x2-y2), 2))
}