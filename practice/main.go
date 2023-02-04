package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs")
	shekar := Employee{"Shekar Singh", 1566}
	fmt.Println(shekar)
	fmt.Printf("%+v\n", shekar)
	fmt.Printf("Name: %v\nEmployee Id : %v\n", shekar.Name, shekar.EmpId)
}

type Employee struct {
	Name  string
	EmpId int
}
