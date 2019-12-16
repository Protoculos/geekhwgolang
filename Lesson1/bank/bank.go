//3 Пользователь вводит сумму вклада в банк и годовой процент.
package main

import "fmt"

func percentFunc(s, per float64) float64 {
	digPer := (s * per) / 100
	s += digPer
	return s
}
func lost(a ...int) []int {
	return int...
}

func main() {
	var score float64
	var percent float64

	fmt.Print("Введите количество денег: ")
	fmt.Scan(&score)
	cicle := score

	fmt.Print("Введите банковский процент: ")
	fmt.Scan(&percent)
	score = percentFunc(score, percent)
	score = percentFunc(score, percent)
	score = percentFunc(score, percent)
	score = percentFunc(score, percent)
	score = percentFunc(score, percent)
	fmt.Println(score)
	//Здесь происходит округление в большую сторону
	fmt.Printf("За пять лет банковский процент по вкладу составит\n%.2f\n", score)
	//Или циклом
	for i := 0; i < 5; i++ {
		cicle = percentFunc(cicle, percent)
	}
	fmt.Println(cicle)
}
