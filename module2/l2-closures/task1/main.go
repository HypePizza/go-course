package main

import "fmt"

func add(n int) func(int) int {
	return func(value int) int {
		return n + value
	}
}

func mult(n int) func(int) int {
	return func(value int) int {
		return n * value
	}
}

func twice(value int, action func(int) int) int {
	value = action(value)
	return action(value)
}

func main() {
	addTwo := add(2)
	multTwo := func(value int) int {
		return 2 * value
	}
	fmt.Println(twice(3, addTwo))
	fmt.Println(twice(3, multTwo))
}
