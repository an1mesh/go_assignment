package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Shape interface {
	Area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}

func (c *Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{0, 0, 5}

	areaOfCircleStruct := c.Area()
	fmt.Println(areaOfCircleStruct)

	totalAreaInterface := totalArea()
	fmt.Println(totalAreaInterface)
}
