package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	value1 := acceptFloatValue(*reader, "value 1")
	value2 := acceptFloatValue(*reader, "value 2")
	fmt.Printf(" value1 : %v , value2: %v \n", value1, value2)
	fmt.Printf("Select operator (+,-,*,/) ")
	operator, _ := acceptOperator(*reader)
	var result float64
	switch strings.TrimSpace(operator) {
	case "+":
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "*":
		result = value1 * value2
	case "/":
		result = value1 / value2
	default:
		fmt.Println("invalid operand !!")
	}

	fmt.Printf("result is : %.2f", result)

}

func acceptFloatValue(reader bufio.Reader, prompt string) float64 {
	fmt.Printf("%v :", prompt)
	input, _ := reader.ReadString('\n')
	inputConvValue, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message := fmt.Sprintf("%v must be a number", prompt)
		panic(message)
	}
	return inputConvValue
}

func acceptOperator(reader bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	fmt.Println("input operand :", input)
	return input, err
}
