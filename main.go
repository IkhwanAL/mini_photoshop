package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ikhwanal/pixel_art_scaler/page"
)

func noCache(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		h.ServeHTTP(w, r)
	})
}

func main() {
	assetsFs := http.FileServer(http.Dir("assets"))
	http.Handle("GET /assets/", http.StripPrefix("/assets/", noCache(assetsFs)))

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		main := page.WebHtml()

		if err := main.Render(r.Context(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, fmt.Errorf("feature not implemented yet").Error(), http.StatusTeapot)
	})

	log.Println("listening on http://localhost:8080")
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatal(err)
	}
}
