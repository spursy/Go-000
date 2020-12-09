package main

import (
	"fmt"
	"net/http"
)

// HelloHandler ..
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8000", nil)
}