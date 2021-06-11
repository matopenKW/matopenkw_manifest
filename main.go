package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test World")
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	fmt.Println("Start Server")

	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":8080", nil)
}
