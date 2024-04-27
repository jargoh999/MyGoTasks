package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("enter numbers to check")
	var input string
	fmt.Scanln(&input)
	if !isPalindrome(input) {
		fmt.Println("this is not palindrome")
	}
	if isPalindrome(input) {
		fmt.Println("this is a palindrome")
	}
}

func palindrome(palindrome string) string {
	checker := ""
	bits := strings.Split(palindrome, "")
	for bit := len(bits) - 1; bit >= 0; bit-- {
		checker += bits[bit]
	}
	return checker
}

func isPalindrome(input string) bool {
	return input == palindrome(input)
}
