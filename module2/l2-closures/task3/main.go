package main

import (
	"fmt"
)

func countMatching(rule func(float64) bool, readings ...float64) int {
	match := 0
	for _, v := range readings {
		if rule(v) {
			match++
		}
	}
	return match
}

func over(limit float64) func(float64) bool {
	return func(value float64) bool {
		return value > limit
	}
}

func inRange(low, high float64) func(float64) bool {
	return func(value float64) bool {
		return value >= low && value <= high
	}
}

func main() {

	var args = []float64{-2, 0, 4.5, 5, 5, 8}
	fmt.Println(countMatching(over(5), args...))
	fmt.Println(countMatching(inRange(0, 5), args...))
	fmt.Println(countMatching(inRange(5, 5), args...))
	fmt.Println(countMatching(over(10), args...))
	fmt.Println(countMatching(inRange(0, 5)))

	fmt.Println(countMatching(func(value float64) bool {
		return value < 0
	}, args...))
}
