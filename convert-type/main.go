package main

import (
	"fmt"
)

func main() {
	age := 18
	marsAge := float64(age)
	fmt.Println(marsAge)

	marsDays := 687.0
	earthDays := 365.2425
	marsAge = marsAge * earthDays / marsDays
	fmt.Println("I am", marsAge, "years old on Mars.")
}