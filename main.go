package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {

		fmt.Println("Error... ", err)

	}
}
