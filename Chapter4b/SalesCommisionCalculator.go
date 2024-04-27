package main

import (
	"fmt"
	"math"
)

func main() {
	totalItems := 0.0
	basePay := 200.0
	for items := 0; items < math.MaxInt; items++ {
		fmt.Println("enter an item")
		var item float64
		fmt.Scanln(&item)
		totalItems += item
		fmt.Println("do you want to add another item")
		var promptAgain string
		fmt.Scanln(&promptAgain)
		if promptAgain == "no" {
			break
		}
	}
	fmt.Println("your wage for the week is ", 0.9*totalItems+basePay)
}
