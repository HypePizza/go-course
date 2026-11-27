package main

import (
	"fmt"
	"strconv"
)

func count() func(int) int {
	sum := 0
	return func(value int) int {
		sum += value
		return sum
	}
}

func markEnd() func(int) int {
	sum := 0
	return func(code int) int {
		if code == 1 {
			return sum
		}
		sum += 1
		return sum
	}

}

func meter(str string, count func(int) int, markEnd func(int) int) (err error) {
	defer markEnd(0)
	value, err := strconv.Atoi(str)
	if err != nil {
		return fmt.Errorf("ошибка: %v", err)
	}
	count(value)
	return err

}

func showRes(count func(int) int, markEnd func(int) int, counterErr int) {
	fmt.Printf("сумма:%d\t", count(0))
	fmt.Printf("завершено попыток:%d\t", markEnd(1))
	fmt.Printf("получено ошибок:%d\n", counterErr)

}

func main() {
	counter := count()
	counterEnd := markEnd()
	counterErr := 0
	err := meter("5", counter, counterEnd)
	if err != nil {
		counterErr += 1
	}
	showRes(counter, counterEnd, counterErr)
	err = meter("abc", counter, counterEnd)
	if err != nil {
		counterErr += 1
	}
	showRes(counter, counterEnd, counterErr)
	err = meter("-2", counter, counterEnd)
	if err != nil {
		counterErr += 1
	}
	showRes(counter, counterEnd, counterErr)
}
