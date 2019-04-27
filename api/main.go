package main

import (
	"log"
	"net/http"

	"github.com/antonmoiseev/now-go-modules/api/handlers"
)

func main() {
	http.HandleFunc("/health", handlers.Health)
	http.HandleFunc("/child", handlers.Child)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
