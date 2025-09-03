// Package routes http request modules
package routes

import "net/http"

func noCache(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		h.ServeHTTP(w, r)
	})
}

func ServerRoute() *http.ServeMux {
	mux := http.NewServeMux()

	assetsFs := http.FileServer(http.Dir("assets"))
	
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", noCache(assetsFs)))

	mux.HandleFunc("GET /", ServerMainPage)

	mux.HandleFunc("POST /upload-image", UploadImage)

	return mux
}
