// MIT License

// Copyright (c) 2017 Alex Ellis
// Copyright (c) 2017 Isaac "Ike" Arias

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package gpio

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// OUTPUT is ...
const OUTPUT = 1

type gpioPin struct {
	valueFd     *os.File
	directionFd *os.File
}

var gpioPins map[string]gpioPin

// Setup does basic initialization
func Setup() {
	gpioPins = make(map[string]gpioPin)
}

// Cleanup does finalization
func Cleanup() {
	for pinNum, v := range gpioPins {
		val, _ := strconv.Atoi(pinNum)
		log.Println("Cleaning up pin ", pinNum)
		v.directionFd.Close()
		v.valueFd.Close()
		unexport(val)
	}
}

func export(pin int) {
	path := "/sys/class/gpio/export"
	bytesToWrite := []byte(strconv.Itoa(pin))
	writeErr := ioutil.WriteFile(path, bytesToWrite, 0644)
	if writeErr != nil {
		log.Panic(writeErr)
	}
}

func unexport(pin int) {
	path := "/sys/class/gpio/unexport"
	bytesToWrite := []byte(strconv.Itoa(pin))
	writeErr := ioutil.WriteFile(path, bytesToWrite, 0644)
	if writeErr != nil {
		log.Panic(writeErr)
	}
}

func pinExported(pin int) bool {
	pinPath := fmt.Sprintf("/sys/class/gpio/gpio%d", pin)
	if file, err := os.Stat(pinPath); err == nil && len(file.Name()) > 0 {
		return true
	}
	return false
}

// PinMode sets the mode in which the Pin will operate
func PinMode(pin int, val int) {
	pinName := strconv.Itoa(pin)

	exported := pinExported(pin)
	if val == OUTPUT {
		if !exported {
			export(pin)
		}
	} else {
		if exported {
			unexport(pin)
		}
	}

	_, exists := gpioPins[pinName]
	if !exists {
		pinPath := fmt.Sprintf("/sys/class/gpio/gpio%d", pin)
		valueFd, openErr := os.OpenFile(pinPath+"/value", os.O_WRONLY, 0640)
		if openErr != nil {
			log.Panic(openErr, pinPath)
		}
		directionFd, openErr := os.OpenFile(pinPath+"/direction", os.O_WRONLY, 0640)
		if openErr != nil {
			log.Panic(openErr, pinPath)
		}
		gpioPins[pinName] = gpioPin{
			valueFd:     valueFd,
			directionFd: directionFd,
		}
		if val == OUTPUT {
			pinDigitalWrite(pin, "out", "direction")
		}
	}
}

// DigitalWrite write a byte to a pin
func DigitalWrite(pin int, val int) {
	pinDigitalWrite(pin, strconv.Itoa(val), "value")
}

func pinDigitalWrite(pin int, val string, mode string) {
	pinName := strconv.Itoa(pin)
	var err error
	if mode == "direction" {
		_, err = gpioPins[pinName].directionFd.Write([]byte(val))
	} else {
		_, err = gpioPins[pinName].valueFd.Write([]byte(val))
	}

	if err != nil {
		log.Panic(err, fmt.Sprintf("Pin: %s Mode: %s Value: %s ", pinName, val, mode))
	}
}