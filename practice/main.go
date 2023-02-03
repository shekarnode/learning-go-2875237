package main

import (
	"fmt"
	"time"
)

func main() {

	n := time.Now()
	fmt.Println("I recorded this video at", n)
	t := time.Date(2000, time.August, 14, 8, 30, 0, 0, time.Local)
	fmt.Println("my birthday is at :", t)

}
