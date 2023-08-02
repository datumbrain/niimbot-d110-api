package main

import (
	"fmt"
	"github.com/datumbrain/label-printer/tag"
	"image"
	"image/png"
	"log"
	"os"
	"time"
)

const printerMacAddress = "08:13:F4:C4:34:53"

func main() {
	err := PrintTag("DB23LPTP3", "https://example.org/%s")
	if err != nil {
		log.Fatalln(err)
	}
}

func PrintTag(text, qrLinkFormat string) error {
	tg := tag.NewGenerator(90, 220)

	img, err := tg.GenerateImage(text, qrLinkFormat)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%d.png", time.Now().UnixMicro())

	err = saveImageToPng(filename, img)
	if err != nil {
		return err
	}

	return runPythonScript("./niimprint", "niimprint/__main__.py", "-a", printerMacAddress, filename)
}

func saveImageToPng(filename string, img image.Image) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}
