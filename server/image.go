package server

import (
	"bytes"
	"image"
	"image/jpeg"
)

// Helper functions for image encoding/decoding (assuming JPEG format)
func decodeImage(data []byte) (image.Image, error) {
	return jpeg.Decode(bytes.NewReader(data))
}

func encodeImage(img image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, nil); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
