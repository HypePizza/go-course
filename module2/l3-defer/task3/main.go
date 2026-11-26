package main

import (
	"errors"
	"fmt"
)

func workOK() error {
	return nil
}

func noData() error {
	return errors.New("нет данных")
}

func clean() error {
	defer fmt.Println("очистка C")
	panic("сбой")
}

func runStep(name string, work func() error) (err error) {
	defer func() {
		if reason := recover(); reason != nil {
			err = fmt.Errorf("%s: паника: %s", name, reason)
		}
	}()
	err = work()
	return err
}

func checkErr(name string, err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s OK\n", name)
}

func main() {
	name := "A"
	err := runStep(name, workOK)
	checkErr(name, err)
	name = "B"
	err = runStep("B", noData)
	checkErr(name, err)
	name = "C"
	err = runStep("C", clean)
	checkErr(name, err)
	name = "D"
	err = runStep(name, workOK)
	checkErr(name, err)
}
