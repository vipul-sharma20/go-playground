package main

import (
	"fmt"
	"math"
)

const val string = "constant string"

func main() {
	fmt.Println(val)

	const num = 10
	const num1 = 4e10 / num

	fmt.Println(float64(num1))
	fmt.Println(math.Sin(num1))
}
