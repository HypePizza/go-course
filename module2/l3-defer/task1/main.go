package main

import "fmt"

func trace(stop bool) {
	n := 2
	//defer fmt.Println("A", n)
	defer func() { fmt.Println("A", n) }()
	n++
	defer func() { fmt.Println("B", n) }()
	if stop {
		return
	}
	n *= 2
	defer fmt.Println("C", n)
}

func main() {
	trace(false)
	//trace(true)
}
