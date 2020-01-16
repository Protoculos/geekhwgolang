package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func parse_args() (file, pat string) {
	if len(os.Args) < 3 {
		log.Fatal("usage: petergrep <file_name> <pattern>")
	}
	file = os.Args[1]
	pat = os.Args[2]
	return
}

func grepFile(file string, pat []byte) string {
	patCount := int64(0)
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if bytes.Contains(scanner.Bytes(), pat) {
			patCount++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	str := "Совпадение '" + string(pat) + "' в файле " + file + " - " + strconv.Itoa(int(patCount))
	return str
}

func main() {
	file, pat := parse_args()
	fmt.Println(grepFile(file, []byte(pat)))
}

// go run grep.go file.txt надо
// Совпадение 'надо' в файле file.txt - 1
