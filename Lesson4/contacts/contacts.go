//2 реализовать для него интерфейс Sort{}
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	pers1 := Person{
		Name:  "Mike",
		Phone: "+17839812348",
	}
	pers2 := Person{
		Name:  "Alice",
		Phone: "+17839749348",
	}
	pers3 := Person{
		Name:  "Xong",
		Phone: "+6723769834",
	}
	var book []Person = []Person{pers1, pers2, pers3}
	fmt.Println(book)
	sort.SliceStable(book, func(i, j int) bool { return book[i].Name < book[j].Name })
	fmt.Println("By name:", book)

	fmt.Println("By name:", book)
}
