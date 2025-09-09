package src

import (
	"net/http"

	"github.com/ikhwanal/pixel_art_scaler/src/database"
	imagecontrol "github.com/ikhwanal/pixel_art_scaler/src/pages/image_control"
)

type Server struct {
	queries *database.Queries
}

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

func (s *Server) ServerRoute() *http.ServeMux {
	mux := http.NewServeMux()

	assetsFs := http.FileServer(http.Dir("assets"))
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", noCache(assetsFs)))

	uploadFs := http.FileServer(http.Dir("uploads"))
	mux.Handle("GET /uploads/", http.StripPrefix("/uploads/", noCache(uploadFs)))
	
	mux.HandleFunc("GET /", serverMainPage)

	imagecontrol.Routes(mux, s.queries)
	return mux
}

func NewServer(queries *database.Queries) *Server {
	return &Server{
		queries: queries,
	}
}
