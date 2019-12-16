//1. Написать функцию, которая определяет, четное ли число.
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

	//У четного числа остаток от деления на 2 будет 0.
	if digit%2 == 0 {
		fmt.Println("Вы ввели четное число.")
	} else {
		fmt.Println("Вы ввели не четное число.")
	}
}
