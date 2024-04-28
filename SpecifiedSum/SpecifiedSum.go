package main

import (
	"fmt"
)

func main() {
	fmt.Println("enter a limit")
	var limit int
	fmt.Scanln(&limit)
	sum := 0
	for counter := 0; counter <= limit; counter++ {
		fmt.Println("enter a number")
		var number int
		fmt.Scanln(&number)
		sum += number

		if sum == limit {
			break
		}
	}
}
