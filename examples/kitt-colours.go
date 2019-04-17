package main

import (
	"fmt"
	"time"

	"github.com/adrianh-za/blinkt-rpi"
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
	var spacing = 360.0 / 16.0
	var hue = 0

	for {

		fmt.Println("pixels: ", firstPixelPosition, secondPixelPosition, thirdPixelPosition, forthPixelPosition)

		var now = time.Now()
		var secs = now.Unix()
		hue = int(secs * 100) % 360
		fmt.Println("hue: ", hue)

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
			var offset = int(3 * spacing)
			fmt.Print("offset: ", offset, " ")
			var pixelHue = ((hue * offset) % 360.0)
			fmt.Println(" pixelHue: ", pixelHue, " ")
			r, g, b := Hsv2Rgb(float64(pixelHue), 1.0, 1.0)
			fmt.Println(" r,g,b: ", r, g, b)
			blinkt.SetPixel(forthPixelPosition, int(r), int(g), int(b))
		}
		//Set third pixel
		if (thirdPixelPosition >= 0) && (thirdPixelPosition <= 7) {
			var offset = int(2 * spacing)
			//fmt.Print("offset: ", offset, " ")
			var pixelHue = ((hue * offset) % 360)
			//fmt.Print("pixelHue: ", pixelHue, " ")
			r, g, b := Hsv2Rgb(float64(pixelHue), 1.0, 1.0)
			//fmt.Println("r, g, b: ", r, g, b)
			blinkt.SetPixel(thirdPixelPosition, int(r), int(g), int(b))
		}
		//Set second pixel
		if (secondPixelPosition >= 0) && (secondPixelPosition <= 7) {
			var offset = int(1 * spacing)
			//fmt.Print("offset: ", offset, " ")
			var pixelHue = ((hue * offset) % 360)
			//fmt.Print("pixelHue: ", pixelHue, " ")
			r, g, b := Hsv2Rgb(float64(pixelHue), 1.0, 1.0)
			//fmt.Println("r, g, b: ", r, g, b)
			blinkt.SetPixel(secondPixelPosition, int(r), int(g), int(b))
		}
		//Set first pixel
		if (firstPixelPosition >= 0) && (firstPixelPosition <= 7) {
			var offset = int(0 * spacing)
			//fmt.Print("offset: ", offset, " ")
			var pixelHue = ((hue * offset) % 360)
			//fmt.Print("pixelHue: ", pixelHue, " ")
			r, g, b := Hsv2Rgb(float64(pixelHue), 1.0, 1.0)
			//fmt.Println("r, g, b: ", r, g, b)
			blinkt.SetPixel(firstPixelPosition, int(r), int(g), int(b))
		}

		//Show and do a small delay
		blinkt.Show()
		time.Sleep(100 * time.Millisecond)
	}
}

func Hsv2Rgb(h, s, v float64) (uint32, uint32, uint32) {
	h /= 360
	if s == 0.0 {
		return uint32(v * 255), uint32(v * 255), uint32(v * 255)
	}
	i := int(h * 6.0)
	f := (h * 6.0) - float64(i)
	p := v * (1.0 - s)
	q := v * (1.0 - s*f)
	t := v * (1.0 - s*(1.0-f))
	i %= 6
	switch i {
	case 0:
		return uint32(v * 255), uint32(t * 255), uint32(p * 255)
	case 1:
		return uint32(q * 255), uint32(v * 255), uint32(p * 255)
	case 2:
		return uint32(p * 255), uint32(v * 255), uint32(t * 255)
	case 3:
		return uint32(p * 255), uint32(q * 255), uint32(v * 255)
	case 4:
		return uint32(t * 255), uint32(p * 255), uint32(v * 255)
	case 5:
		return uint32(v * 255), uint32(p * 255), uint32(q * 255)
	default:
		return 0, 0, 0
	}
}
