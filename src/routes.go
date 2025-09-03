package src

import (
	"net/http"

	imagecontrol "github.com/ikhwanal/pixel_art_scaler/src/pages/image_control"
)

func serverMainPage(w http.ResponseWriter, r *http.Request) {
	main := WebHtml()

	if err := main.Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

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

	mux.HandleFunc("GET /", serverMainPage)

	mux.HandleFunc("POST /upload-image", imagecontrol.UploadImage)

	return mux
}
