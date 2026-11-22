package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)

func fail(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "ошибка: "+format+"\n", args...)
	os.Exit(1)
}

func humanize(value int) (string, error) {
	if value < 0 {
		return "", fmt.Errorf("размер не может быть отрицательным")
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
		return strconv.Itoa(value) + " " + unit, nil
	}
	return strconv.FormatFloat(float64(value)/divisor, 'f', 2, 64) + " " + unit, nil
}

func toBytes(value float64, unit string) (int64, error) {
	if value < 0 {
		return 0, fmt.Errorf("размер не может быть отрицательным")
	}
	var multiplier float64
	switch strings.ToUpper(unit) {
	case "PB":
		multiplier = PB
	case "TB":
		multiplier = TB
	case "GB":
		multiplier = GB
	case "MB":
		multiplier = MB
	case "KB":
		multiplier = KB
	case "B":
		multiplier = B
	default:
		return 0, fmt.Errorf("неизвестная единица %q, нужна B, KB, MB, GB, TB или PB", unit)
	}
	return int64(value * multiplier), nil
}

func main() {
	args := os.Args
	mode := len(args) - 1

	if mode < 1 || mode > 2 {
		fmt.Fprintf(os.Stderr, "использование:\n  go run . <байты>              → в удобные единицы\n  go run . <число> <единица>    → в байты\n")
		os.Exit(1)
	}
	switch mode {
	case 1:
		value, err := strconv.Atoi(args[1])
		if err != nil {
			fail("%q — не целое число байт", args[1])
		}
		res, err := humanize(value)

		fmt.Println(res)

	case 2:
		value, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			fail("%q - не число байт", args[1])
		}

		res, err := toBytes(value, args[2])
		if err != nil {
			fail("%s", err)
		}
		fmt.Printf("%d B\n", res)
	}
}
