package imagecontrol

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"

	"github.com/ikhwanal/pixel_art_scaler/src/database"
)

func Routes(mux *http.ServeMux, queries *database.Queries) {
	mux.HandleFunc("POST /upload-image", uploadImageHadler(queries))
}

type FileHeader struct {
	Filename string
	Width    int
	Height   int
	Size     int
}

func uploadImageHadler(q *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
			Width:    img.Bounds().Dx(),
			Height:   img.Bounds().Dy(),
			Size:     int(header.Size),
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

		id, err := q.UploadImage(r.Context(), database.UploadImageParams{
			Filename:         header.Filename,
			CurrentVersionID: 1,
		})
		if err != nil {
			http.Error(w, "unable to load an image", http.StatusInternalServerError)
			return
		}

		_, err = q.AddImageToTrack(r.Context(), database.AddImageToTrackParams{
			ImageID:         id,
			ParentVersionID: 0,
			Filename:        header.Filename,
			Operation:       "Init",
		})
		if err != nil {
			http.Error(w, "unable to load an image", http.StatusInternalServerError)
			return
		}

		w.Header().Set(
			"HX-Trigger",
			fmt.Sprintf("{\"updateCanvas\": {\"image\": \"%s\", \"width\": %d, \"height\": %d, \"image_id\": %d}}", uploadPath, img.Bounds().Dx(), img.Bounds().Dy(), id),
		)

		err = EditOption(fileHeader, id).Render(r.Context(), w)
		if err != nil {
			http.Error(w, "failed to show result", http.StatusInternalServerError)
			return
		}
	}
}

func applyScale(q *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
