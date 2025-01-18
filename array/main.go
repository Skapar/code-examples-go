package main

import (
	"fmt"
)

func main() {
	var planets [8]string
	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"
	
	earth := planets[2]
	fmt.Println(earth)

	fmt.Println(len(planets))

	fmt.Println(planets[3] == "")

	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	for i := 0 ; i < len(dwarfs); i++ {
		fmt.Println(dwarfs[i])
	}

	planetsSecond := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	for i := 0; i < len(planetsSecond); i++ {
		fmt.Println(planetsSecond[i])
	}

	planetsMarkII := planetsSecond

	planets[2] = "упс"

	fmt.Println(planetsSecond[2])
	fmt.Println(planetsMarkII)

	TwoDimesionalArray()
}