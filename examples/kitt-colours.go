package main

import (
	"fmt"
	"time"

	"github.com/adrianh-za/blinkt-rpi"
	"github.com/adrianh-za/utils-golang/colorsys"
)

func main() {
	brightness := 0.2
	blinkt := blinkt.NewBlinkt(brightness)
	blinkt.ClearOnExit = true
	blinkt.CaptureExit = true
	blinkt.Setup()

	//Start positions of pixels
	var firstPixelPosition = 0
	var secondPixelPosition = -1
	var thirdPixelPosition = -2
	var forthPixelPosition = -3
	var pixelDirection = 1
	var hue = 0.0

	for {
		fmt.Println("hue-pixels/: ", hue, firstPixelPosition, secondPixelPosition, thirdPixelPosition, forthPixelPosition)

		//Set pixels positions
		firstPixelPosition = firstPixelPosition + pixelDirection
		secondPixelPosition = firstPixelPosition - pixelDirection
		thirdPixelPosition = firstPixelPosition - (pixelDirection * 2)
		forthPixelPosition = firstPixelPosition - (pixelDirection * 3)

		//Check for change of direction
		if firstPixelPosition == 8 {
			pixelDirection = -1
			firstPixelPosition = 6
		} else if firstPixelPosition == -1 {
			pixelDirection = 1
			firstPixelPosition = 1
		}

		//Switch off all LEDs
		blinkt.Clear()

		//Set fourth pixel
		if (forthPixelPosition >= 0) && (forthPixelPosition <= 7) {
			r, g, b := colorsys.Hsv2Rgb(hue, 1.0, 0.1)	//Lower the VALUE passed in, lower the brightness
			blinkt.SetPixel(forthPixelPosition, int(r), int(g), int(b))
		}
		//Set third pixel
		if (thirdPixelPosition >= 0) && (thirdPixelPosition <= 7) {
			r, g, b := colorsys.Hsv2Rgb(hue, 1.0, 0.25)  //Lower the VALUE passed in, lower the brightness
			blinkt.SetPixel(thirdPixelPosition, int(r), int(g), int(b))
		}
		//Set second pixel
		if (secondPixelPosition >= 0) && (secondPixelPosition <= 7) {
			r, g, b := colorsys.Hsv2Rgb(hue, 1.0, 0.5)  //Lower the VALUE passed in, lower the brightness
			blinkt.SetPixel(secondPixelPosition, int(r), int(g), int(b))
		}
		//Set first pixel
		if (firstPixelPosition >= 0) && (firstPixelPosition <= 7) {
			r, g, b := colorsys.Hsv2Rgb(hue, 1.0, 1.0)  //Lower the VALUE passed in, lower the brightness
			blinkt.SetPixel(firstPixelPosition, int(r), int(g), int(b))
		}

		//Show and do a small delay
		blinkt.Show()
		time.Sleep(100 * time.Millisecond)

		//Increment the hue up to max of 360, then reset
		hue = hue + 5
		if (hue > 360) {
			hue = 0
		}
	}
}