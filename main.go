package main

import (
	"log"
	"net/http"

	"github.com/antonmoiseev/now-go-modules/api"
)

func main() {
	http.HandleFunc("/health", api.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
