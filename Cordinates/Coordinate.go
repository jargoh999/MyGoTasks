package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Print("Enter x-coordinate of Point 1: ")
	fmt.Scan(&x1)
	fmt.Print("Enter y-coordinate of Point 1: ")
	fmt.Scan(&y1)
	fmt.Print("Enter x-coordinate of Point 2: ")
	fmt.Scan(&x2)
	fmt.Print("Enter y-coordinate of Point 2: ")
	fmt.Scan(&y2)
	if isPerpendicular(x1, y1, x2, y2) {
		fmt.Println("The points are perpendicular")
	}
	if !isPerpendicular(x1, y1, x2, y2) {
		fmt.Println("The points are not perpendicular")
	}
	fmt.Println("The distance between the points is", calculateDistance(x1, y1, x2, y2))
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func isPerpendicular(x1, y1, x2, y2 float64) bool {
	return x1 == y1 || y2 == x2
}
