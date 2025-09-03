package main

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/ikhwanal/pixel_art_scaler/routes"
)

func main() {
	log.Println("listening on http://localhost:8080")
	if err := http.ListenAndServe("127.0.0.1:8080", routes.ServerRoute()); err != nil {
		log.Fatal(err)
	}
}
