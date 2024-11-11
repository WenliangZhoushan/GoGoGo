package main

import (
	"fmt"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	content, _ := os.ReadFile("./hello.txt")
	_, _ = fmt.Fprintln(w, string(content))
}

func main() {
	http.HandleFunc("/hello", sayHello)
	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
