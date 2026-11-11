package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})
	log.Println("сервер запущен: http://localhost:8080/health")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
