package main

import (
	"fmt"
	"time"

	"github.com/adrianh-za/blinkt-rpi"
	"github.com/adrianh-za/utils-golang/colorsys"
)

func main() {
	brightness := 0.1
	blinkt := blinkt.NewBlinkt(brightness)
	blinkt.ClearOnExit = true
	blinkt.CaptureExit = true
	blinkt.Setup()

	for {
		//Get the seconds since Epoch (divide NanoSeconds by 1e6 to get seconds)
		var seconds = time.Now().UnixNano() / 1e6
		
		//Calculate the base hue from the seconds
		var hue = seconds % 360

		//Calculate hue for each pixel on Blinkt
		for pixelCount := 0; pixelCount <= 7; pixelCount++ {
			var offset = pixelCount * 30
			var pixelHue = (hue + int64(offset)) % 360
			var r, g, b = colorsys.Hsv2Rgb(float64(pixelHue), 1.0, 1.0)
			blinkt.SetPixel(pixelCount, int(r), int(g), int(b))
			fmt.Println("hue: ", hue, " pixelCount: ", pixelCount, " pixelHue", pixelHue, "RGB", r,"|",g,"|",b)
		}

		//Show and do a small delay
		blinkt.Show()
		time.Sleep(10 * time.Millisecond)
	}
}