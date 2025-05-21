package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
