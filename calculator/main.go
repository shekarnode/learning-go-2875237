package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("A simple calculator")
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("value 1:")
	input1, _ := reader.ReadString('\n')
	fmt.Printf("value 2:")
	input2, _ := reader.ReadString('\n')
	aFloatValue1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	aFloatValue2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)

	if err != nil {
		panic(err)
	}

	aFloatSum := aFloatValue1 + aFloatValue2
	fmt.Printf("Sum : %.2f \n ", aFloatSum)

}
