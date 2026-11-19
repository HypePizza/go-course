package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "использование go run . <число>")
		os.Exit(1)
	}
	num, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ошибка: %q - не целое число\n", args[1])
		os.Exit(1)
	}
	if num < 2 {
		fmt.Fprintln(os.Stderr, "ошибка: число меньше 2. нужно ввести число больше или равное 2")
		os.Exit(1)
	}

	var divisors []int
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
		}
	}

	if len(divisors) == 2 {
		fmt.Printf("%d - простое\n", num)
		return
	}
	fmt.Printf("%d - составное, делители:", num)
	for _, d := range divisors {
		fmt.Printf(" %d", d)
	}
	fmt.Println()
}
