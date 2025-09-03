package imagecontrol

import (
	"io"
	"net/http"
	"os"
	"strconv"
)

type FileHeader struct {
	Filename string
	Width int
	Height int
	Size int
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
		http.Error(w, "opps failed to get width file header", http.StatusInternalServerError)
		return
	} 
	height, err := strconv.Atoi(r.FormValue("height"))
	if err != nil {
		http.Error(w, "oops failed to get height file header", http.StatusInternalServerError)
		return
	}

	fileHeader := FileHeader{
		Filename: header.Filename,
		Width: width,
		Height: height,
		Size: int(header.Size),
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
	err = OverViewImage(fileHeader).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "failed to show result", http.StatusInternalServerError)
	}
}
