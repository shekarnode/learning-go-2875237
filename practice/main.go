package main

import (
	"fmt"
)

func main() {
	value := 24.5
	var p = &value
	fmt.Println("value of p is :", *p)

	value2 := "shekar"
	pointer1 := &value2
	fmt.Println("address of value 2: ", &value2)
	fmt.Println("value of value2 : ", value2)
	fmt.Println("address of pointer1 : ", pointer1)
	fmt.Println("value of pointer1 : ", *pointer1)

	//changing value through pointer

	value3 := "singh"
	fmt.Println("value of value3 : ", value3)
	pointer2 := &value3
	fmt.Println("value of pointer2 : ", *pointer2)
	*pointer2 = "shekar singh"
	fmt.Println("new value of value3 :", value3)

}
