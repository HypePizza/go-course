package main

import "fmt"

func main() {
	for n := 1; n <= 30; n++ {
		switch {
		case n%3 == 0 && n%5 == 0:
			fmt.Println("FizzBuzz")
		case n%3 == 0:
			fmt.Println("Fizz")
		case n%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(n)
		}
	}
}
