package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	port := 3333

	fmt.Println("Server is running on port", port)

	err := http.ListenAndServe(":"+fmt.Sprint(port), nil)
	if err != nil {
		fmt.Println(err)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello")
}
