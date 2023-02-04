package main

import (
	"fmt"
)

func main() {
	theAnswer := 42
	if theAnswer < 0 {
		fmt.Println("Less than Zero")
	} else if theAnswer == 0 {
		fmt.Println("Equals to zero")
	} else {
		fmt.Println("Greater than zero")
	}

	if x := -42; x < 0 {
		fmt.Println("Less than Zero")
	} else if x == 0 {
		fmt.Println("Equals to zero")
	} else {
		fmt.Println("Greater than zero")
	}
}
