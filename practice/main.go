package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	// dow := rand.Intn(7) + 1
	// fmt.Println("Day", dow)

	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		fmt.Println("it's sunday !!")
		// fallthrough
	case 2:
		fmt.Println("it's monday !!")
	default:
		fmt.Println("it's some other day!!")
	}
}
