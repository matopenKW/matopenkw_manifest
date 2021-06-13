package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test World")
	fmt.Fprintf(w, "Hello, World")
}

func hostname(w http.ResponseWriter, r *http.Request) {
	hn := os.Getenv("HOSTNAME")
	fmt.Println(hn)
	fmt.Fprintf(w, hn)
}

func main() {
	fmt.Println("Start Server")

	http.HandleFunc("/", handler)
	http.HandleFunc("/hostname", hostname)
	http.ListenAndServe(":8080", nil)
}
