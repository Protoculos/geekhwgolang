package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/*Пакет ioutil предоставляет функцию для чтения каждого байта в файле и возврата его в виде байтового среза. Эта функция удобна тем, что вам не нужно определять срез байта перед выполнением чтения. Недостатком является то, что действительно большой файл будет возвращать большой фрагмент, который может быть больше, чем ожидалось.*/
func main() {
	// Open file for reading
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	//Делаем сразу вывод в string
	fmt.Printf("Data as string: %s\n", data)

}
