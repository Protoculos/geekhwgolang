//4. * Заполнить массив из 100 элементов различными простыми числами.
package main

import "fmt"

// return list of primes less than N
func sieveOfEratosthenes(N int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		if len(primes) == 100 {
			break
		}
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return
}

func main() {
	primes := sieveOfEratosthenes(600)
	for _, p := range primes {
		fmt.Println(p)
	}
	fmt.Println("Количество простых чисел: ", len(primes))
}
