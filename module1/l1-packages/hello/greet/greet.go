package greet

import "fmt"

// Hello начинается с заглавной буквы — видна из других пакетов
func Hello(name string) string {
	return fmt.Sprintf("Привет, %s!", decorate(name))
}

// decorate начинается со строчной буквы — приватная функция пакета
func decorate(name string) string {
	return "*" + name + "*"
}
