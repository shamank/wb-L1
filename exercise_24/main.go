package main

import (
	"fmt"
	"math"
	"wb-l1/exercise_24/point"
)

//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

func main() {
	a := point.NewPoint(10, 10)
	b := point.NewPoint(15, 10)

	d := DistanceCalc(a, b)
	fmt.Printf("Distance beetwen points: %v", d)

}

func DistanceCalc(p1 *point.Point, p2 *point.Point) float64 {
	coords1 := p1.GetCoords()
	coords2 := p2.GetCoords()

	dx := coords2[0] - coords1[0]
	dy := coords2[1] - coords1[1]

	return math.Sqrt(dx*dx + dy*dy)
}
