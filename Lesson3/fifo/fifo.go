//3. * Реализовать очередь.
package main

import "fmt"

func main() {
	var newDigit, lastDigit int
	fifo := []int{1, 2, 3}
	route := "да"
	for route == "да" {
		fmt.Print("Введите простое число: ")
		fmt.Scan(&newDigit)
		lastDigit = fifo[0]
		fifo = append(fifo[1:], newDigit)
		fmt.Println(lastDigit, fifo)
		fmt.Print("Если необходимо продолжить введите 'да': ")
		fmt.Scan(&route)
	}

}
