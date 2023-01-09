package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Print("GO API STARTING")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n", r)
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n", r)
		t := time.Now().Format(time.RFC1123)
		fmt.Fprintf(w, t)
	})

	http.ListenAndServe(":8000", nil)
}
