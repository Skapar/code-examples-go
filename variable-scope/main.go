package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	switch month := rand.Intn(12) + 1; month { // переменные era, year и month в области видимости    
    case 2:
        day := rand.Intn(28) + 1 // новый день                 
        fmt.Println(era, year, month, day)
    case 4, 6, 9, 11:
        day := rand.Intn(30) + 1                  
        fmt.Println(era, year, month, day)
    default:
        day := rand.Intn(31) + 1                   
        fmt.Println(era, year, month, day)
    } // month и day за пределами области видимости    
}