package main

import (
	"fmt"
	"time"
)

func Switch() {
	fmt.Println("Here entry to cave and path to east")
	command := "come inside"
	
	switch command {
	case "go east":
		fmt.Println("You are going to mountains.")
	case "come inside", "enter cave":
		fmt.Println("You are in a cave.")
	case "read sign":
		fmt.Println("Sign in outside wrote: 'Minors not allowed'.")
	default:
		fmt.Println("You are standing still.")	
	}

	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("Today is Monday.")
	case time.Tuesday:
		fmt.Println("Today is Tuesday.")
	case time.Wednesday:
		fmt.Println("Today is Wednesday.")
	case time.Thursday:
		fmt.Println("Today is Thursday.")
	case time.Friday:
		fmt.Println("Today is Friday.")
	case time.Saturday:
		fmt.Println("Today is Saturday.")
	case time.Sunday:
		fmt.Println("Today is Sunday.")
	}

	Repeat()

	switch time.Now().Weekday() {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("It's a weekday.")
	default:
		fmt.Println("It's a weekend.")
	}

	Repeat()

	size := "XXXL"

	switch size {

	case "XXS":
		fmt.Println("Extra Extra Small")
	case "XS":
		fmt.Println("Extra Small")
	case "S":
		fmt.Println("Small")
	case "M":
		fmt.Println("Medium")
	case "L":
		fmt.Println("Large")
	case "XL":
		fmt.Println("Extra Large")
	case "XXL":
		fmt.Println("Extra Extra Large")
	case "XXXL":
		fmt.Println("Extra Extra Extra Large")
	default:
		fmt.Println("Unknown size")
	}

	Repeat()

	switch num := 6; num % 2 == 0 {
	case true:
		fmt.Println(num, "is even")
	case false:
		fmt.Println(num, "is odd")
	}
}