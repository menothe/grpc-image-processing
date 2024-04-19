package server

import (
	"image"
	"image/color"
)

// converts image to grayscale
func convertImageToGrayscale(img image.Image, workerPool chan func()) image.Image {

	// Convert each pixel to grayscale
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := uint8((float64(r) * 0.299) + (float64(g) * 0.587) + (float64(b) * 0.114))
			grayImg.Set(x, y, color.Gray{Y: gray})
		}
	}
	return grayImg
}
