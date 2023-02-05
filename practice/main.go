package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 10, "woof!!"}
	poodleptr := &poodle
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
	poodle.speak()
	poodle.Sound = "Arf !!"
	poodle.speak()
	poodleptr.speakThreeTimes()
	poodleptr.speakThreeTimes()
}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) speak() {
	fmt.Printf(d.Sound)
}

func (d Dog) speakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}
