package main

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/ikhwanal/pixel_art_scaler/src"
)

func main() {
	log.Println("listening on http://localhost:8080")
	if err := http.ListenAndServe("127.0.0.1:8080", src.ServerRoute()); err != nil {
		log.Fatal(err)
	}
}
