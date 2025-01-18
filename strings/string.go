package main

import "fmt"

func striiing() {
	message := "shalom"
	c := message[5]
	fmt.Println(c)

	for i := 0; i < len(message); i++ {
		c := message[i]
		fmt.Printf("%c\n", c)
	}
}