package imagecontrol

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"net/http"
	"os"
)

type FileHeader struct {
	Filename string
	Width int
	Height int
	Size int
}

func UploadImage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // Max Size
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

	img, formatImage, err := image.Decode(file)
	if err != nil {
		http.Error(w, "oops failed to decode given image", http.StatusBadRequest)
		return
	}

	fileHeader := FileHeader{
		Filename: header.Filename,
		Width: img.Bounds().Dx(),
		Height: img.Bounds().Dy(),
		Size: int(header.Size),
	}

	uploadPath := "./uploads/" + header.Filename
	outFile, err := os.Create(uploadPath)
	if err != nil {
		http.Error(w, "unable to save file", http.StatusInternalServerError)
		return
	}

	if formatImage == "png" {
		err = png.Encode(outFile, img)
	} else {
		err = jpeg.Encode(outFile, img, nil)
	}
	if err != nil {
		http.Error(w, "failed to write file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", fmt.Sprintf("{\"updateCanvas\": {\"image\": \"%s\"}}", uploadPath))
	err = OverViewImage(fileHeader).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "failed to show result", http.StatusInternalServerError)
		return
	}
}
