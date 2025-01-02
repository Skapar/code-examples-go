package main

import (
	"fmt"
)

func calculator() {
	var firstInput, secondInput int
	var operation string

	fmt.Println("Enter a first number: ")
	fmt.Scan(&firstInput)

	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scan(&operation)

	fmt.Println("Enter a second number: ")
	fmt.Scan(&secondInput)

	switch operation {
	case "+":
		fmt.Println(firstInput + secondInput)
	case "-":
		fmt.Println(firstInput - secondInput)
	case "*":
		fmt.Println(firstInput * secondInput)
	case "/":
		fmt.Println(firstInput / secondInput)
	default:
		fmt.Println("Invalid operation")
	}
}