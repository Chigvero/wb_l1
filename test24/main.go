package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func distance(p1 *Point, p2 *Point) float64 {
	return math.Sqrt((math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2)))
}
func main() {
	p1 := Point{3, 4}
	p2 := Point{0, 0}
	fmt.Print(distance(&p1, &p2))
}
