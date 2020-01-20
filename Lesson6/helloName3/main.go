//Task 3
//Program can fetch int and string.
//Name http://localhost:4000/?hello&name=MyName
//Int http://localhost:4000/?hello&name=456
//Можно вставить парсинг и других имен. Здесь решил сделать валидацию числа и строки
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("name"))
	if err != nil || id < 1 {
		word := r.URL.Query().Get("name")
		if word == "" {
			fmt.Println("Enter the name")
			return
		}
		fmt.Fprintf(w, "Specific name is %s.", word)
		return

	}

	fmt.Fprintf(w, "Display a specific ID %d...", id)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatalln("Server not starting: ", err)

}
