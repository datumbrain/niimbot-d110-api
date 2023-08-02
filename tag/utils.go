package tag

import (
	"golang.org/x/image/draw"
	"image"
	"image/color"
)

func joinImages(outputHeight, outputWidth int, image1, image2 image.Image) *image.RGBA {
	// Create a new RGBA image with the desired output size
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

func rotateImageClockwise(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	rotatedImage := image.NewRGBA(image.Rect(0, 0, height, width))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			newX := height - y - 1
			newY := x
			rotatedImage.Set(newX, newY, img.At(x, y))
		}
	}

	return rotatedImage
}

func rotateImageCounterClockwise(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	rotatedImage := image.NewRGBA(image.Rect(0, 0, height, width))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			newX := y
			newY := width - x - 1
			rotatedImage.Set(newX, newY, img.At(x, y))
		}
	}

	return rotatedImage
}
