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
