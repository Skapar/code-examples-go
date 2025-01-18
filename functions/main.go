package main

import (
	"fmt"
)

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celciusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

func kelvinToFahrenheit(k float64) float64 {
	c := kelvinToCelsius(k)
	return celciusToFahrenheit(c)
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Println(kelvin, "K is", celsius, "C")

	fahrenheit := celciusToFahrenheit(celsius)
	fmt.Println(celsius, "C is", fahrenheit, "F")

	fahrenheit = kelvinToFahrenheit(kelvin)
	fmt.Println(kelvin, "K is", fahrenheit, "F")
}