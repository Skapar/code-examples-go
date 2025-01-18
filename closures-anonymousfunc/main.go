package main

import (
	"fmt"
)

var v = func(message string) {
	fmt.Println(message)
}

func main() {
	v("Hello, World!")

	func() {
		fmt.Println("Hello, World!")
	}()

	f := func(message string) {
		fmt.Println(message)
	}

	f("Hello, World!")
}