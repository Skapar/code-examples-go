package main

import (
	"fmt"
	"math/rand"
)

func loop() {
	count := 10

	for count > 0 {
		fmt.Println(count)
		// time.Sleep(time.Second)
		count--
	}
	fmt.Println("Launch!")

	Repeat()

	degrees := 0

	for {
		fmt.Print(degrees, " ")

		degrees++
		if degrees >= 360 {
			degrees = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}
}