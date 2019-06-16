package main

import (
	"fmt"
	"time"

	"github.com/adrianh-za/go-blinkt-rpi"
)

func main() {
	brightness := 0.5
	blinkt := blinkt.NewBlinkt(brightness)
	blinkt.ClearOnExit = true
	blinkt.CaptureExit = true
	blinkt.Setup()

	time.Sleep(100 * time.Millisecond)

	//Start positions of pixels
	var firstPixelPosition = 0
	var secondPixelPosition = -1
	var thirdPixelPosition = -2
	var forthPixelPosition = -3
	var pixelDirection = 1

	for {

		fmt.Println("pixels: ", firstPixelPosition, secondPixelPosition, thirdPixelPosition, forthPixelPosition)

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

		blinkt.Clear()

		//Set fourth pixel
		if (forthPixelPosition >= 0) && (forthPixelPosition <= 7) {
			blinkt.SetPixel(forthPixelPosition, 8, 0, 0)
		}
		//Set third pixel
		if (thirdPixelPosition >= 0) && (thirdPixelPosition <= 7) {
			blinkt.SetPixel(thirdPixelPosition, 32, 0, 0)
		}
		//Set second pixel
		if (secondPixelPosition >= 0) && (secondPixelPosition <= 7) {
			blinkt.SetPixel(secondPixelPosition, 96, 0, 0)
		}
		//Set first pixel
		if (firstPixelPosition >= 0) && (firstPixelPosition <= 7) {
			blinkt.SetPixel(firstPixelPosition, 255, 0, 0)
		}

		//Show and do a small delay
		blinkt.Show()
		time.Sleep(100 * time.Millisecond)
	}
}
