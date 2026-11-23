package main

import (
	"fmt"
	"os"
	"strconv"
)

func parseNumbers(args ...string) ([]float64, error) {
	if len(args) < 1 {
		return []float64{}, fmt.Errorf("ошибка: нужно хотя бы одно число")
	}

	var res []float64
	for _, str := range args {
		value, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return []float64{}, fmt.Errorf("ошибка: %q - не число", str)
		}
		res = append(res, value)
	}
	return res, nil
}

// нет проверки на пустой слайс. Отсечен функцией parseNumbers
func stats(nums ...float64) (sum, avg, lo, hi float64) {
	for index, value := range nums {
		if index == 0 {
			lo = value
			hi = value
		}
		sum += value
		if value < lo {
			lo = value
		}
		if value > hi {
			hi = value
		}
	}
	avg = sum / float64(len(nums))
	return sum, avg, lo, hi
}

func main() {
	args := os.Args
	values, err := parseNumbers(args[1:]...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	sum, avg, min, max := stats(values...)
	fmt.Printf("количество: %d\n", len(values))
	fmt.Printf("сумма: %.2f\n", sum)
	fmt.Printf("среднее: %.2f\n", avg)
	fmt.Printf("мин: %.2f, макс: %.2f\n", min, max)
}
