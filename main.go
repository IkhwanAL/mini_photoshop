package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ikhwanal/pixel_art_scaler/src"
	"github.com/ikhwanal/pixel_art_scaler/src/database"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "mini_photoshop.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := database.New(db)

	server := src.NewServer(queries)

	log.Println("listening on http://localhost:8080")
	if err := http.ListenAndServe("127.0.0.1:8080", server.ServerRoute()); err != nil {
		log.Fatal(err)
	}
}
