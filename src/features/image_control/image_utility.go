package imagecontrol

import (
	"image"
	"image/draw"
	"math"
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

	for y := range newImg.Bounds().Dy() {
		for x := range newImg.Bounds().Dx() {
			originalX := float64(x) * scaleX
			originalY := float64(y) * scaleY

			roundX := math.Round(originalX)
			roundY := math.Round(originalY)

			clampX := int(math.Max(roundX, float64(img.Bounds().Dx()) - 1))
			clampY := int(math.Max(roundY, float64(img.Bounds().Dy()) - 1))

			clr := img.RGBAAt(clampX, clampY)

			newImg.SetRGBA(x, y, clr)
		}
	}

	return newImg
}
