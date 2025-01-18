package main

import "fmt"

func main() {
	peace := "peace"
	var peaceTwo = "peace"
	var peaceThird string = "peace"

	fmt.Println(peace, peaceTwo, peaceThird)

	fmt.Println("peace be upon you\nupon you be peace")
	fmt.Println(`strings can span multiple lines with the \n escape sequence`)

	fmt.Println(`
    peace be upon you
    upon you be peace`)

	fmt.Printf("%v is a %[1]T\n", "literal string")
	fmt.Printf("%v is a %[1]T\n", `raw string literal`)

	runeAndBytes()
	striiing()

}