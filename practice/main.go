package main

import (
	"fmt"
)

func main() {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
	fmt.Println("----------------------------------------------------------------")
	for i := range colors {
		fmt.Println(colors[i])
	}

	for i, color := range colors {
		fmt.Printf("%v : %v \n", i, color)
	}

	//labels
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("sum :", sum)
		if sum > 200 {
			goto theEnd
		}
	}

theEnd:
	fmt.Println("End of Program")
}
