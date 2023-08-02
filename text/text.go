package text

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"os"
)

type Hinting string

const (
	Full Hinting = "full"
	None Hinting = "none"
)

type Config struct {
	Height       int     // image height in points
	Width        int     // image width in points
	DPI          float64 // resolution in Dots Per Inch
	Padding      int     // text left and right padding
	FontFile     string  // path of the font to use
	font         *truetype.Font
	FontSize     float64 // font size in points
	CharsPerLine int     // chars displayed per line
	Hinting      Hinting // none | full
	Spacing      float64 // line spacing
	WhiteOnBlack bool    // white text on a black background
}

func GetImage(c Config, text string) (*image.RGBA, error) {
	var err error
	c.font, err = readFontFile(c.FontFile)
	if err != nil {
		return nil, err
	}

	if c.CharsPerLine > 0 {
		c.FontSize = float64(c.Width-c.Padding*2) / float64(c.CharsPerLine) * 72 / c.DPI
	}

	fg, bg := image.Black, image.White
	if c.WhiteOnBlack {
		fg, bg = bg, fg
	}

	rgba := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))
	draw.Draw(rgba, rgba.Bounds(), bg, image.Point{}, draw.Src)

	// Freetype context
	ctx := freetype.NewContext()
	ctx.SetDPI(c.DPI)
	ctx.SetClip(rgba.Bounds())
	ctx.SetDst(rgba)
	ctx.SetFont(c.font)
	ctx.SetFontSize(c.FontSize)
	ctx.SetSrc(fg)

	switch c.Hinting {
	case Full:
		ctx.SetHinting(font.HintingFull)
	default:
		ctx.SetHinting(font.HintingNone)
	}

	face := truetype.NewFace(c.font, &truetype.Options{
		Size: c.FontSize,
		DPI:  c.DPI,
	})

	// Calculate the widths and print to image
	pt := freetype.Pt(c.Padding, ctx.PointToFixed(c.FontSize).Round())
	newline := func() {
		pt.X = fixed.Int26_6(c.Padding) << 6
		pt.Y += ctx.PointToFixed(c.FontSize * c.Spacing)
	}

	for _, x := range text {
		w, _ := face.GlyphAdvance(x)
		if x == '\t' {
			x = ' '
		} else if c.font.Index(x) == 0 {
			continue
		} else if pt.X.Round()+w.Round() > c.Width-c.Padding {
			newline()
		}

		pt, err = ctx.DrawString(string(x), pt)
		if err != nil {
			return nil, err
		}
	}

	return rgba, nil
}

func readFontFile(file string) (*truetype.Font, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return truetype.Parse(b)
}
