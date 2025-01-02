package main

import (
	"fmt"
	"strings"
)

func Cave() {
	fmt.Println("You are in a cave.")

	var command = "leave outside"
	var exit = strings.Contains(command, "leave")

	fmt.Println("You have chosen to", exit)
}