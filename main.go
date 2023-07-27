package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	dpi      = 240.0 // "screen resolution in Dots Per Inch"
	fontfile = "/System/Library/Fonts/Supplemental/Arial.ttf"
	hinting  = "full" // "none | full"
	size     = 6.0    // "font size in points"
	width    = 130    //220    // "image width in points"
	padding  = 10     // "text left and right padding"
	height   = 30     // "image height in points"
	chars    = 0      //  "chars displayed per line"
	spacing  = 1.0    // "line spacing"
	wonb     = false  // "white text on a black background"
)

func main() {
	parseFlags()

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

	for _, x := range []rune("DB23LPTP3") {
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
			log.Fatal(err)
		}
	}

	saveImage(rgba)
}

func parseFlags() {
	flag.Parse()

	if chars > 0 {
		size = float64(width-padding*2) / float64(chars) * 72 / dpi
	}
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

func saveImage(rgba *image.RGBA) {
	rgba = joinQR(rgba)

	f, err := os.OpenFile(fmt.Sprintf("%d.png", time.Now().UnixMicro()), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panic(err)
	}

	bf := bufio.NewWriter(f)
	if err := png.Encode(bf, rgba); err != nil {
		log.Panic(err)
	}
	if err := bf.Flush(); err != nil {
		log.Panic(err)
	}
}

func joinQR(text image.Image) *image.RGBA {
	qr := getQr()

	return joinImages(qr, text)

	//starting position of the second image (bottom left)
	sp2 := image.Point{qr.Bounds().Dx(), 0}

	//new rectangle for the second image
	r2 := image.Rectangle{sp2, sp2.Add(text.Bounds().Size())}

	//rectangle for the big image
	r := image.Rectangle{image.Point{0, 0}, r2.Max}

	rgba := image.NewRGBA(r)

	draw.Draw(rgba, qr.Bounds(), qr, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, r2, text, image.Point{0, 0}, draw.Src)

	return rgba
}

func getQr() image.Image {
	var img []byte
	img, err := qrcode.Encode("https://example.org", qrcode.Medium, 60)
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
