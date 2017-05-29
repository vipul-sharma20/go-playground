package main

import "fmt"

func main() {
	defer fmt.Println("Hey there")

	fmt.Println("I was first")
}
