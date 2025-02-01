package main

import (
	"fmt"
)

func main(){
	// creating map
	temperature := map[string]int {
		"Earth": 15,
		"Mars": -65,
	}
	temp := temperature["Earth"]
	fmt.Printf("On Earth the temperature is %v C\n", temp)

	// adding new key-value pair
	temperature["Earth"] = 16
	temperature["Venus"] = 464

	fmt.Println(temperature)

	moon := temperature["Moon"]
	fmt.Println(moon)

	temperature["Moon"] = -20

	// check if key exists
	if moon, ok := temperature["Moon"]; ok {
		fmt.Println("Ok is", ok)
		fmt.Printf("On Moon the temperature is %v C\n", moon)
	} else {
		fmt.Println("Ok is", ok)
		fmt.Println("Where is the Moon?")
	}

	// copy map
	temperatureSecond := temperature
	fmt.Println(temperatureSecond)

	// make map
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	frequency := make(map[float64]int)

	for _, t := range temperatures {
		frequency[t]++
	}

	for t, num := range frequency {
		fmt.Printf("%v C occurs %v times\n", t, num)
	}

	// delete key-value pair
	delete(temperature, "Earth")
	
	// set map
	var temperaturesSet = []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	
	set := make(map[float64]bool)
	for _, t := range temperaturesSet {
		set[t] = true
	}

	if set[-28.0] {
		fmt.Println("part of set")
	}



}