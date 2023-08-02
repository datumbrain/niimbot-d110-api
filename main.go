package main

import (
	"fmt"
	"image"
	"image/png"
	"img/tag"
	"log"
	"os"
	"time"
)

func main() {
	tg := tag.NewGenerator(90, 220)

	img, err := tg.GenerateImage("DB23LPTP3", "https://example.org/%s")
	if err != nil {
		log.Fatalln(err)
	}

	filename := fmt.Sprintf("%d.png", time.Now().UnixMicro())

	err = saveImageToPng(filename, img)
	if err != nil {
		log.Fatalln(err)
	}

	//mac := "08:13:F4:C4:34:53"
	//niimprintScript := "niimprint/__main__.py"
	//
	//exec.Command("python3", niimprintScript, "-a", mac, name)
}

func saveImageToPng(filename string, img image.Image) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}
