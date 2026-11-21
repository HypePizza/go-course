package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	B  = iota + 1
	KB = 1 << (iota * 10)
	MB
	GB
	TB
	PB
)

func main() {
	args := os.Args
	mode := len(args) - 1

	if mode < 1 || mode > 2 {
		fmt.Fprintf(os.Stderr, "использование:\n  go run . <байты>              → в удобные единицы\n  go run . <число> <единица>    → в байты")
		os.Exit(1)
	}

	switch mode {
	case 1:
		value, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "ошибка: %q — не целое число байт\n", args[1])
			os.Exit(1)
		}
		if value < 0 {
			fmt.Fprintln(os.Stderr, "ошибка: размер не может быть отрицательным")
			os.Exit(1)
		}
		var unit string
		var divisor float64
		switch {
		case value >= PB:
			unit = "PB"
			divisor = PB
		case value >= TB:
			unit = "TB"
			divisor = TB
		case value >= GB:
			unit = "GB"
			divisor = GB
		case value >= MB:
			unit = "MB"
			divisor = MB
		case value >= KB:
			unit = "KB"
			divisor = KB
		default:
			unit = "B"
			divisor = 1
		}
		if unit == "B" {
			fmt.Printf("%d %s\n", value, unit)
		} else {
			fmt.Printf("%.2f %s\n", float64(value)/divisor, unit)
		}

	case 2:
		value, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ошибка: %q - не число байт\n", args[1])
			os.Exit(1)
		}
		if value < 0 {
			fmt.Fprintln(os.Stderr, "ошибка: размер не может быть отрицательным")
			os.Exit(1)
		}
		unit := strings.ToUpper(args[2])
		if unit != "B" && unit != "KB" && unit != "MB" &&
			unit != "GB" && unit != "TB" && unit != "PB" {
			fmt.Fprintf(os.Stderr, "ошибка: неизвестная единица %q, нужна B, KB, MB, GB, TB или PB\n", args[2])
			os.Exit(1)
		}
		switch unit {
		case "PB":
			fmt.Printf("%d B\n", int64(value*PB))
		case "TB":
			fmt.Printf("%d B\n", int64(value*TB))
		case "GB":
			fmt.Printf("%d B\n", int64(value*GB))
		case "MB":
			fmt.Printf("%d B\n", int64(value*MB))
		case "KB":
			fmt.Printf("%d B\n", int64(value*KB))
		case "B":
			fmt.Printf("%d B\n", int64(value))
		}
	}
}
