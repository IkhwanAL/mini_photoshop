package imagecontrol

import (
	"image"
	"image/draw"
)

func toRGBA(img image.Image) *image.RGBA {
	newImage := image.NewRGBA(img.Bounds())

	draw.Draw(newImage, img.Bounds(), img, img.Bounds().Min, draw.Src)

	return newImage
}

func nearestNeightbor(img *image.RGBA, width int, height int) *image.RGBA {
	scaleX := float64(img.Bounds().Dx()) / float64(width)
	scaleY := float64(img.Bounds().Dy()) / float64(height)

	newImg := image.NewRGBA(img.Bounds())

	for y := range img.Bounds().Dy() {
		for x := range img.Bounds().Dx() {
			originalX := float64(x) * scaleX
			originalY := float64(y) * scaleY


		}
	}

	return img
}
