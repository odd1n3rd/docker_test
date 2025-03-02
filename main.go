package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Docker!")
		fmt.Fprintln(w, "Now alpine version!")
	})

	http.ListenAndServe(":12345", nil)
}
