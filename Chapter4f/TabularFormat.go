package main

import "fmt"

func main() {
	fmt.Println("N   N2   N3   N4 ")
	for number := 1; number <= 5; number++ {
		fmt.Println(number, "  ", number*number, " ", number*number*number, " ", number*number*number*number)
	}

}
