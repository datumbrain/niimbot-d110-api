package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"
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

	name := fmt.Sprintf("%d.png", time.Now().UnixMicro())

	img, err := generateTagImage("DB23LPTP3", "https://example.org/%s")
	if err != nil {
		log.Fatalln(err)
	}

	err = saveImageToPng(name, img)
	if err != nil {
		log.Fatalln(err)
	}

	mac := "08:13:F4:C4:34:53"
	niimprintScript := "niimprint/__main__.py"

	exec.Command("python3", niimprintScript, "-a", mac, name)
}

func parseFlags() {
	flag.Parse()

	if chars > 0 {
		size = float64(width-padding*2) / float64(chars) * 72 / dpi
	}

	switch runtime.GOOS {
	case "linux":
		fontfile = "/usr/share/fonts/Arial.ttf"
	case "darwin":
		fontfile = "/System/Library/Fonts/Supplemental/Arial.ttf"
	case "windows":
		fontfile = `C:\Windows\Fonts\Arial.ttf`
	}

}
