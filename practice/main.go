package main

import "fmt"

func main() {
	var colors [2]string
	colors[0] = "red"
	colors[1] = "pink"

	fmt.Println("values of colors array is :", colors)
	fmt.Println("value of 0th index in colors array is: ", colors[0])
	fmt.Println("length of colors array is :", len(colors))
	fmt.Println("---------------------------------------------")
	//single line declaration and initialization of array
	number := [3]int{1, 2, 3}
	fmt.Println("values in number array is :", number)
	fmt.Println("length of number array is :", len(number))

}
