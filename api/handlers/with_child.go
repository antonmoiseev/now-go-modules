package handlers

import (
	"fmt"
	"net/http"

	"github.com/antonmoiseev/now-go-modules/handlers/child"
)

func Child(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, child.Name())
}
