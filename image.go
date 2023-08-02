package main

import (
	"bytes"
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/skip2/go-qrcode"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func generateTagImage(tag string, qrLinkFormat string) (*image.RGBA, error) {

	text, err := getTextImage(tag)
	if err != nil {
		return nil, err
	}

	qr := getQrImage(fmt.Sprintf(qrLinkFormat, tag))

	return joinImages(qr, text), nil
}

func getTextImage(tag string) (*image.RGBA, error) {
	f := readFont()

	fg, bg := image.Black, image.White
	if wonb {
		fg, bg = bg, fg
	}

	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(rgba, rgba.Bounds(), bg, image.Point{}, draw.Src)

	// Freetype context
	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetFont(f)
	c.SetFontSize(size)
	c.SetSrc(fg)
	switch hinting {
	default:
		c.SetHinting(font.HintingNone)
	case "full":
		c.SetHinting(font.HintingFull)
	}

	opts := truetype.Options{}
	opts.Size = size
	opts.DPI = dpi
	face := truetype.NewFace(f, &opts)

	// Calculate the widths and print to image
	pt := freetype.Pt(padding, c.PointToFixed(size).Round())
	newline := func() {
		pt.X = fixed.Int26_6(padding) << 6
		pt.Y += c.PointToFixed(size * spacing)
	}

	var err error

	for _, x := range tag {
		w, _ := face.GlyphAdvance(x)
		if x == '\t' {
			x = ' '
		} else if f.Index(x) == 0 {
			continue
		} else if pt.X.Round()+w.Round() > width-padding {
			newline()
		}

		pt, err = c.DrawString(string(x), pt)
		if err != nil {
			return nil, err
		}
	}

	return rgba, nil
}

func readFont() *truetype.Font {
	b, err := os.ReadFile(fontfile)
	if err != nil {
		log.Panic(err)
	}
	f, err := truetype.Parse(b)
	if err != nil {
		log.Panic(err)
	}

	return f
}

func getQrImage(data string) image.Image {
	var img []byte
	img, err := qrcode.Encode(data, qrcode.Medium, 60)
	if err != nil {
		log.Panic(err)
	}

	imgx, err := png.Decode(bytes.NewReader(img))
	if err != nil {
		log.Panic(err)
	}

	return imgx
}

func joinImages(image1, image2 image.Image) *image.RGBA {
	// Create a new RGBA image with the desired output size (220x90)
	outputWidth := 220
	outputHeight := 90
	output := image.NewRGBA(image.Rect(0, 0, outputWidth, outputHeight))

	// Fill the entire output image with white color
	draw.Draw(output, output.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)

	// Calculate the starting X position to center the two images
	startX := (outputWidth - image1.Bounds().Dx() - image2.Bounds().Dx()) / 2

	// Calculate the starting Y position to vertically center the images
	startY := (outputHeight - image1.Bounds().Dy()) / 2

	// Draw the first image onto the output image
	draw.Draw(output, image.Rect(startX, startY, startX+image1.Bounds().Dx(), startY+image1.Bounds().Dy()), image1, image.Point{}, draw.Over)

	// Calculate the starting X position for the second image
	startX += image1.Bounds().Dx()

	// Calculate the starting Y position for the second image
	startY = (outputHeight - image2.Bounds().Dy()) / 2

	// Draw the second image onto the output image
	draw.Draw(output, image.Rect(startX, startY, startX+image2.Bounds().Dx(), startY+image2.Bounds().Dy()), image2, image.Point{}, draw.Over)

	return output
}
