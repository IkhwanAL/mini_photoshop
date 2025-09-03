package routes

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/ikhwanal/pixel_art_scaler/page"
)

type FileHeader struct {
	Filename string
	Width int
	Height int
}

func UploadImage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "failed to parse multipart form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("uploadedFile")
	if err != nil {
		http.Error(w, "file not found", http.StatusBadRequest)
		return
	}
	defer file.Close()

	width, err := strconv.Atoi(r.FormValue("width"))
	if err != nil {
		http.Error(w, "opps my mistake need to handle the error", http.StatusInternalServerError)
	} 
	height, err := strconv.Atoi(r.FormValue("height"))

	fileHeader := FileHeader{
		Filename: header.Filename,
		Width: width,
		Height: height,
	}

	outFile, err :=os.Create("./uploads/" + header.Filename)
	if err != nil {
		http.Error(w, "unable to save file", http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, "failed to write file", http.StatusInternalServerError)
		return
	}
	page.OverViewImage().Render(r.Context(), w)

	return
}
