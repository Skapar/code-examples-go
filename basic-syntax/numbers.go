package main

import (
	"fmt"
	"math/rand"
)

func GetNumbers() {
	num := rand.Intn(10) + 10
	fmt.Println(num)

	numTwo := rand.Intn(10) + 1

	fmt.Println(num * numTwo)
	fmt.Println(num / numTwo)
	fmt.Println(num % numTwo)
	fmt.Println(num - numTwo)
	fmt.Println(num + numTwo)
}