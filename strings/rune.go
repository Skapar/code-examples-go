package main

import "fmt"

func runeAndBytes() {
	// type byte = uint8
	// type rune = int32

	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	
	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)

	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang)

	grade := 'A'
	var gradeOne = 'A'
	var gradeTwo rune = 'A'

	fmt.Printf("%v %v %v\n", grade, gradeOne, gradeTwo)

	var star byte = '*'
	fmt.Printf("%c %[1]v\n", star)

}