package main

import (
	"bufio"
	"image"
	"image/png"
	"os"
)

func saveImageToPng(filename string, img *image.RGBA) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(bufio.NewWriter(f), img)
}
