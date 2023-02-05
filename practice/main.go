package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions")
	doSomething()
	sum := addValues(4, 5)
	fmt.Println("sum is :", sum)

	multiSum, multiCount := multiValues(5, 2, 6, 7)
	fmt.Println("multisum value is :", multiSum)
	fmt.Println("multisum count is : ", multiCount)
}

func doSomething() {
	fmt.Println("Doing something")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}

func multiValues(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}
