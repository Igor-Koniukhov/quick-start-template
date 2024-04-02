package main

import (
	"fmt"
	"net/http"
)

func HandleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	http.HandleFunc("/", HandleTest)
	http.ListenAndServe(":8080", nil)
}
