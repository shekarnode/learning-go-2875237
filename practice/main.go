package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	colors = append(colors, "orange")
	fmt.Println("new color slice is :", colors)
	fmt.Println("---------------------------------------")
	fmt.Println("---Removing values from colors slice---")
	colors = append(colors[0 : len(colors)-1])
	fmt.Println("new color slice is :", colors)
	fmt.Println("-----------------------------------------------")
	fmt.Println("---creating new slice object via make() func---")
	numbers := make([]int, 3, 3)
	numbers[0] = 55
	numbers[1] = 100
	numbers[2] = 22
	fmt.Println("numbers slice is :", numbers)
	numbers = append(numbers, 52)
	fmt.Println("new numbers slice is: ", numbers)
	fmt.Println("-----------------------------------")
	fmt.Println("-----------sorting slice-----------")
	sort.Ints(numbers)
	fmt.Println("sorted numbers slice is: ", numbers)
	sort.Strings(colors)
	fmt.Println("sorted colors slice is: ", colors)
}
