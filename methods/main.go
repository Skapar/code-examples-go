package main

import (
	"fmt"
)

type celsius float64
type kelvin float64

func kelvinToCelsius(k kelvin) celsius {
    return celsius(k - 273.15) // Необходима конвертация типа
}

func (c celsius) celciusToKelvin() kelvin {
	return kelvin(c + 273.15)
}

func main() {
	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	fmt.Println(k, "° K is", c, "° C")

	var c1 celsius = 127.0
	k1 := c1.celciusToKelvin()
	fmt.Println(c1, "° C is", k1, "° K")
}