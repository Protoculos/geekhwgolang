package main

import (
	"flag"
	"io"
	"log"
	"os"
)

func main() {

	of := flag.String("file", "file.txt", "An origin file")
	nf := flag.String("nfile", "newfile.txt", "A new file")

	flag.Parse()
	// Open original file
	originalFile, err := os.Open(*of)
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	// Create new file
	newFile, err := os.Create(*nf)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// Copy the bytes to destination from source
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)

	// Commit the file contents
	// Flushes memory to disk
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
// go run copyfile.go -file file.txt -nfile newfile.txt
// 2020/01/16 13:59:10 Copied 56 bytes.
