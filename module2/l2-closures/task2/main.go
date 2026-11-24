package main

import "fmt"

func newTotal(start int) func(int) int {
	return func(value int) int {
		start += value
		return start
	}
}

func main() {
	a := newTotal(10)
	b := newTotal(100)
	fmt.Println(a(3))
	fmt.Println(a(-2))
	fmt.Println(b(5))
	fmt.Println(a(0))
	c := newTotal(11)
	fmt.Println(c(4))
	fmt.Println(a(0))
	fmt.Println(c(-1))
	fmt.Println(b(0))
}
