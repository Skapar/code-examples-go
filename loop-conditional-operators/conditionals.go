package main

import (
	"fmt"
)

func Conditionals() {
	fmt.Println("Sign in outside wrote: 'Minors not allowed'.")

	age := 17
	adult := age >= 18

	fmt.Printf("In age %v, am i an adult? %v.\n", age, adult)
}