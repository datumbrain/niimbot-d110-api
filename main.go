package main

import (
	"fmt"
	"github.com/datumbrain/label-printer/API"
	"github.com/datumbrain/label-printer/tag"
	"github.com/gin-gonic/gin"
	"image"
	"image/png"
	"net/http"
	"os"
	"time"
)

const printerMacAddress = "08:13:F4:C4:34:53"
const fixedTagText = "DB23LPTP3"
const fixedQRLinkFormat = "https://example.org/%s"

func main() {
	r := gin.Default()
	r.POST("/print", handlePrintRequest)
	r.Run(":8080")
}

func PrintTag(text, qrLinkFormat string) error {
	tg := tag.NewGenerator(96, 220)

	img, err := tg.GenerateImage(text, qrLinkFormat)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%d.png", time.Now().UnixMicro())

	err = saveImageToPng(filename, img)
	if err != nil {
		return err
	}

	return runPythonScript(".", "./niimprint/niimprint/__main__.py", "-a", printerMacAddress, filename)
}

func saveImageToPng(filename string, img image.Image) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}

func handlePrintRequest(c *gin.Context) {

	err := PrintTag(fixedTagText, fixedQRLinkFormat)
	if err != nil {
		API.SendJSONResponse(c, gin.H{"error": "Error printing"}, http.StatusInternalServerError)
		return
	}

	API.SendJSONResponse(c, gin.H{"message": "Print request sent successfully"}, http.StatusOK)
}
