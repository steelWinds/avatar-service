package main

import (
	"log"
	"net/http"

	"github.com/steelWinds/identavatar/internal/app"
)

func main() {
	http.Handle("/", http.HandlerFunc(app.HandlerGetIdentcoin))

	err := http.ListenAndServe(":3180", nil)

	if err != nil {
		log.Fatal("Server shutdown:", err)
	}
}
