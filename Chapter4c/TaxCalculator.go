package main

import "fmt"

func main() {
	totalTax := 0.0
	for citizen := 1; citizen <= 3; citizen++ {
		fmt.Println("enter your name")
		var name string
		fmt.Scanln(&name)
		fmt.Println("enter your earnings")
		var earnings float64
		fmt.Scanln(&earnings)
		if earnings > 30_000 {
			totalTax = 0.2 * earnings
		}
		totalTax = 0.15 * earnings
		fmt.Println("Citizen ", name, "your tax is", totalTax)
	}
}
