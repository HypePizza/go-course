package main

import "fmt"

func main() {

	//Вариант 1
	for i := range 10 {
		v := 5 - i*2 // 5, 3, 1, -1, -3, ...
		switch {
		case v < 0:
			fmt.Println("отрицательное, стоп")
			goto Outerloop
		default:
			fmt.Println(v)
		}
	}
Outerloop:
	//Вариант 2
	for i := range 10 {
		v := 5 - i*2 // 5, 3, 1, -1, -3, ...
		if v < 0 {
			fmt.Println("отрицательное, стоп")
			break
		}
		fmt.Println(v)
	}
}
