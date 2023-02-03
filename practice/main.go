package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("integer sum is :", intSum)
	fmt.Println("Math")

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("float sum is : ", floatSum)
	fmt.Printf("Circumference: %.1f\n", floatSum)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("circumference of circle is : %.2f\n", circumference)
}
