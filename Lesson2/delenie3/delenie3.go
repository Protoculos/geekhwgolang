//2. Написать функцию, которая определяет, делится ли число без остатка на 3.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number string
	var digit int

	//Получаем число
	fmt.Print("Введите число: ")
	fmt.Scanln(&number)

	//Конвертируем стринговую значение в int
	digit, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("Произошла ошибка конвертации")
	}

	//В данном задании по условию остаток от деления на 3 будет 0.
	if digit%3 == 0 {
		fmt.Println("Введенное число делится на 3 без остатка.")
	} else {
		fmt.Println("Введенное число не делится на 3 без остатка.")
	}
}
