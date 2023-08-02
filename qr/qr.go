package qr

import (
	"bytes"
	"github.com/skip2/go-qrcode"
	"image"
	"image/png"
)

func GetImage(data string, size int) (image.Image, error) {
	img, err := qrcode.Encode(data, qrcode.Medium, size)
	if err != nil {
		return nil, err
	}

	return png.Decode(bytes.NewReader(img))
}
