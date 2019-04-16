package main

import (
	"fmt"
	"time"
	"../blinkt"
)

func main() {
	brightness := 0.5
	blinkt := blinkt.NewBlinkt(brightness)
	blinkt.ClearOnExit = true
	blinkt.CaptureExit = true
	blinkt.Setup()


	time.Sleep(100 * time.Millisecond)

	var firstPixelPosition = 0
	var secondPixelPosition = -1
	var thirdPixelPosition = -2
	var forthPixelPosition = -3
	var pixelDirection = 1


	for {

		fmt.Println("pixels: ", firstPixelPosition, secondPixelPosition, thirdPixelPosition, forthPixelPosition)

		firstPixelPosition = firstPixelPosition + pixelDirection
		secondPixelPosition = firstPixelPosition - pixelDirection
		thirdPixelPosition = firstPixelPosition - (pixelDirection * 2)
		forthPixelPosition = firstPixelPosition - (pixelDirection * 3)

		if firstPixelPosition == 8 {
				pixelDirection = -1
				firstPixelPosition = 6
		} else if firstPixelPosition == -1 {
				pixelDirection = 1
				firstPixelPosition = 1
		}

		blinkt.Clear()
		if (forthPixelPosition >= 0) && (forthPixelPosition <= 7) {
			//blinkt.SetPixel(forthPixelPosition, 0, 8, 0)
			blinkt.SetPixel(forthPixelPosition, 8, 0, 0)
		}

		if (thirdPixelPosition >= 0) && (thirdPixelPosition <= 7) {
			//blinkt.SetPixel(thirdPixelPosition, 0, 32, 0)
			blinkt.SetPixel(thirdPixelPosition, 32, 0, 0)
		}

		if (secondPixelPosition >= 0) && (secondPixelPosition <= 7) {
			//blinkt.SetPixel(secondPixelPosition, 0, 96, 0)
			blinkt.SetPixel(secondPixelPosition, 96, 0, 0)
		}

		if (firstPixelPosition >= 0) && (firstPixelPosition <= 7) {
			//blinkt.SetPixel(firstPixelPosition, 0, 255, 0)
			blinkt.SetPixel(firstPixelPosition, 255, 0, 0)
		}
		blinkt.Show()

		time.Sleep(100 * time.Millisecond)
	}	
}
