package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Print("Введите значение для катета a: ")
	fmt.Scan(&a)
	fmt.Print("Введите значение для катета b: ")
	fmt.Scan(&b)

	//Вычисляем площадь треугольника square
	sq := (a * b) / 2
	fmt.Println("Площадь прямоугольного треугольника равна:", sq)

	//Вычисляем гипотенузу треугольника hypotenuse
	hyp := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	fmt.Println("Гипотенуза - ", hyp)
	//Сокращаем вывод до 2х знаков после точки. Вывод будет округлен
	fmt.Printf("Гипотенуза %.2f\n", hyp)

	//Вычисляем периметр треугольника perimeter
	per := a + b + hyp
	fmt.Println("Периметр треугольника равен: ", per)
	//Сокращаем вывод до 2х знаков после точки. Вывод будет округлен
	fmt.Printf("Периметр треугольника равен %.2f\n", per)
}
