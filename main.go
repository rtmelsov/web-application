package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprint(w, "Hello World!")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Number of ytes written: %d\n", n)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
