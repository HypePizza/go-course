package main

import (
	"fmt"
	"os"
	"runtime"

	"go-course/module1/l1-packages/hello/calc"
	"go-course/module1/l1-packages/hello/greet" // свой пакет: имя модуля + путь к папке
)

func main() {
	// os.Args[0] — путь к программе, дальше идут аргументы командной строки
	name := "мир"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Println(greet.Hello(name))
	fmt.Println(calc.Sum(1, 1))
	fmt.Println(runtime.GOOS + "/" + runtime.GOARCH)
}
