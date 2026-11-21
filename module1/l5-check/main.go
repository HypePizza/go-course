package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	_  = iota
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
		fmt.Println("использование:")
		fmt.Println("\tgo run . <байты>\t\t→ в удобные единицы")
		fmt.Println("\tgo run . <число> <единица>\t→ в байты")
		os.Exit(1)
	}

	switch mode {
	case 1:
		value, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "ошибка: %q - не целое число байт\n", args[1])
			os.Exit(1)
		}
		if value < 0 {
			fmt.Fprintln(os.Stderr, "ошибка: размер не может быть отрицательным")
			os.Exit(1)
		}
		switch {
		case value >= PB:
			fmt.Printf("%.2f PB\n", float64(value)/PB)
		case value >= TB:
			fmt.Printf("%.2f TB\n", float64(value)/TB)
		case value >= GB:
			fmt.Printf("%.2f GB\n", float64(value)/GB)
		case value >= MB:
			fmt.Printf("%.2f MB\n", float64(value)/MB)
		case value >= KB:
			fmt.Printf("%.2f KB\n", float64(value)/KB)
		default:
			fmt.Printf("%d B\n", value)
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
			fmt.Fprintf(os.Stderr, "ошибка: неизвестная еденица %q, нужна B, KB, MB, GB, TB или PB\n", unit)
		}
		switch unit {
		case "PB":
			fmt.Printf("%.f B\n", value*PB)
		case "TB":
			fmt.Printf("%.f B\n", value*TB)
		case "GB":
			fmt.Printf("%.f B\n", value*GB)
		case "MB":
			fmt.Printf("%.f B\n", value*MB)
		case "KB":
			fmt.Printf("%.f B\n", value*KB)
		case "B":
			fmt.Printf("%.f B\n", value)
		}
	}
}
