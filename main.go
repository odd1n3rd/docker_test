package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Docker!")
		fmt.Fprintln(w, "Now alpine version!")
		fmt.Fprintln(w, "v1.2, absolutely light-weight cause of multi-stage build!")
	})

	http.ListenAndServe(":12345", nil)
}
