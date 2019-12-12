package main

import (
	"fmt"
	"strconv"
)

const (
	dollRate = 63.22
)

//Функция приводит число к двум знакам после запятой.
func roundNum(n float64) float64 {
	numInt := int(n * 100)
	n = float64(numInt) / 100
	return n
}

func main() {
	var rub float64
	fmt.Print("Введите количество рублей для конвертации: ")
	fmt.Scan(&rub)
	doll := rub / dollRate
	dollRoundNum := roundNum(doll)
	fmt.Println("Вы можете получить за эту сумму ", dollRoundNum, "$")

	//Можно без функции
	//Правда здесь происходит округление, и строго говоря для банка не подойдет.
	fmt.Printf("Вы можете получить за эту сумму %.2f$\n", doll)

	//Впринципе можно получить float с двумя знаками из строковой переменной, но здесь происходит округление.
	//Поэтому для банка не подойдет
	digitString := fmt.Sprintf("%.2f", doll)
	digitFloat, _ := strconv.ParseFloat(digitString, 64)
	//Или одной строкой digitFloat, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", doll), 64)
	fmt.Println("Вы можете получить за эту сумму ", digitFloat, "$")
}
