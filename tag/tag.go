package tag

import (
	"github.com/datumbrain/label-printer/qr"
	"github.com/datumbrain/label-printer/text"
	"image"
)

type Generator struct {
	height int
	width  int
}

func NewGenerator(height, width int) *Generator {
	return &Generator{height: height, width: width}
}

func (g Generator) GenerateImage(tag, qrText string) (image.Image, error) {
	// getting QR image
	qrCode, err := qr.GetImage(qrText, 90)
	if err != nil {
		return nil, err
	}

	// getting text image
	txt, err := text.GetImage(text.Config{
		Height:       25,
		Width:        130,
		DPI:          240.0,
		Padding:      10,
		FontFile:     "fonts/Arial.ttf",
		FontSize:     6.0,
		Hinting:      text.Full,
		Spacing:      1.0,
		WhiteOnBlack: false,
	}, tag)
	if err != nil {
		return nil, err
	}

	// joining and rotating the image
	finalImage := joinImages(g.height, g.width, qrCode, txt)

	return rotateImageCounterClockwise(finalImage), nil
}
