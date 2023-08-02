package utils

import (
	"fmt"
	"github.com/datumbrain/label-printer/tag"
	"time"
)

func PrintTag(text, qrLinkFormat string) error {
	tg := tag.NewGenerator(96, 220)

	img, err := tg.GenerateImage(text, qrLinkFormat)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%d.png", time.Now().UnixMicro())

	err = SaveImageToPng(filename, img)
	if err != nil {
		return err
	}

	return runPythonScript(".", "./niimprint/niimprint/__main__.py", "-a", printerMacAddress, filename)
}
