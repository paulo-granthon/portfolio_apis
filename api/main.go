package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	io.WriteString(w, "Hello World")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
	io.WriteString(w, "Hello")
}
