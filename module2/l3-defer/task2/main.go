package main

import (
	"errors"
	"fmt"
	"strconv"
)

func readSample(raw string, begin func() error, finish func()) (int, error) {
	if err := begin(); err != nil {
		fmt.Println("результат открытия: ошибка")
		return 0, err
	}
	defer finish()
	fmt.Println("рузультат открытия: успех")
	value, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("ошибка: парсинг %q", raw)
	}
	return value, nil

}

func beginOK() error {
	return nil
}

func beginBusy() error {
	return errors.New("ошибка: датчик занят")
}

func finish() {
	fmt.Println("закрытие")
}

func main() {
	fmt.Println(readSample("17", beginOK, finish))
	fmt.Println(readSample("abc", beginOK, finish))
	fmt.Println(readSample("17", beginBusy, finish))
}
