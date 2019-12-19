package main
import (
    "fmt"
    "os"
)

func main() {
		// text := "Hello"
		f, err := os.Create("hello.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
    // file, err := os.Create("hello.txt")

    // if err != nil{
    //     fmt.Println("Unable to create file:", err)
    //     os.Exit(1)
    // }
    // defer file.Close()
    if _, err = f.WriteString("My name is Mike"); err != nil {
			fmt.Println(err)
			}

    fmt.Println("Done.")
}