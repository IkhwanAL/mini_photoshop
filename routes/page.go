package routes

import (
	"net/http"

	"github.com/ikhwanal/pixel_art_scaler/page"
)

func ServerMainPage(w http.ResponseWriter, r *http.Request) {
	main := page.WebHtml()

	if err := main.Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
