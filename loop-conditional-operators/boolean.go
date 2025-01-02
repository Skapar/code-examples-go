package main

import (
	"fmt"
)

func BooleanOp() {
	fmt.Println("In outside 2100 year. Is it a leap year?")

	year := 2100
	leap := year%400 == 0 || year%4 == 0 && year%100 != 0

	if leap {
		fmt.Println("Yes, it is a leap year.")
	} else {
		fmt.Println("No, it is not a leap year.")
	}

	haveTorch := true
	litTorch := false

	if !haveTorch || !litTorch {
		fmt.Println("You can't see anything.")
	}
}