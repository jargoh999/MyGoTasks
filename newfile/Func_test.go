package main

import (
	"fmt"
	"testing"
)

func TestAdder(test *testing.T) {
	value := make([]int, 0)
	for number := 1; number < 10; number++ {
		value = append(value, number)
	}
	fmt.Println(value)
	newValue := adder(value)
	fmt.Println(newValue)
	if len(newValue) != 9 {
		test.Errorf("len(value) = %d; want %d", len(value), 9)
	}

	if newValue[0] != 2 {
		test.Errorf("value[0] = %d; want %d", value[0], 2)
	}

}
