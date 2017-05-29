package main

import "fmt"

func test(n ...int) {
	fmt.Println(n)
}

func main() {
	test(2, 3, 4, 5)

	num := []int{1, 2, 3, 4}
	test(num...)
}
