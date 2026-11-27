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

func markErr() func(err error) int {
	sumErr := 0
	return func(err error) int {
		if err != nil {
			sumErr += 1
			return sumErr
		}
		return sumErr
	}
}

func meter(str string, count func(int) int, markErr func(error) int) (res int, err error) {
	defer func() {
		if err != nil {
			markErr(err)
		}
	}()
	value, err := strconv.Atoi(str)
	if err != nil {
		err = fmt.Errorf("ошибка: %v", err)
	}
	return count(value), err

}

func showRes(count func(int) int, markErr func(error) int) {
	fmt.Printf("сумма:%d\t", count(0))
	fmt.Printf("кол-во ошибок:%d\n", markErr(nil))

}

func main() {
	counter := count()
	counterErr := markErr()
	meter("5", counter, counterErr)
	showRes(counter, counterErr)
	meter("abc", counter, counterErr)
	showRes(counter, counterErr)
	meter("-2", counter, counterErr)
	showRes(counter, counterErr)
}
