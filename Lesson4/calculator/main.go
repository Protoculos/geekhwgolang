//3 калькулятор help
package main

import (
	"fmt"
)

func main() {
	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}
		if input == "exit" {
			break
		}
		if res, err := Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else if hp, err := Help(input); err == nil {
			fmt.Printf("Помощь: %v\n", hp)
		} else {
			fmt.Println("Не удалось произвести вычисление")
		}
	}
}
