package main

import (
	"log"
	"net/http"

	"github.com/antonmoiseev/now-go-modules/handlers"
)

func main() {
	http.HandleFunc("/health", handlers.Health)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
