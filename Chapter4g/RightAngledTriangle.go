package main

import "fmt"

func main() {
	for asterisk := 0; asterisk < 10; asterisk++ {
		for nextLine := 0; nextLine < asterisk; nextLine++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}
}
