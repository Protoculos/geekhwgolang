//4. * Заполнить массив из 100 элементов различными простыми числами.
//Как в задании
package main

import "fmt"

func generator(n int) []int {
	dig := make(map[int]string)
	var count int
	var prime []int
	for i := 2; i <= n; i++ {
		dig[i] = ""
	}

	for i := 2; i <= n; i++ {
		for j := 2; j <= n; j++ {
			if dig[i] != "зачеркнуто" && i <= n {
				count = i * j
				if count <= n {
					dig[count] = "зачеркнуто"
				} else {
					continue
				}
			}
		}
	}

	for k, v := range dig {
		if v != "зачеркнуто" {
			prime = append(prime, k)
			if len(prime) == 100 {
				break
			}
		}
	}
	return prime
}

func main() {
	prime := generator(600)
	for _, v := range prime {
		fmt.Println(v)
	}
	fmt.Println("Длина слайса: ", len(prime))
}
