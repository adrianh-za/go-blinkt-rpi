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

	var hue = 0.0

	for {
		fmt.Println("hue: ", hue)
		
		//Switch off all LEDs
		blinkt.Clear()

		//Calculate RGB from Hue
		r, g, b := colorsys.Hsv2Rgb(hue, 1.0, 1.0)
		blinkt.SetAll(int(r), int(g), int(b))

		//Show and do a small delay
		blinkt.Show()
		time.Sleep(20 * time.Millisecond)

		//Increment the hue up to max of 360, then reset
		hue++
		if (hue > 360) {
			hue = 0
		}
	}
}