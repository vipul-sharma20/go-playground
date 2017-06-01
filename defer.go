package main

import "fmt"

func main() {
	defer fmt.Println("Hey there")

	fmt.Println("I was first")

    i := mythical()
    fmt.Println("return: ", i)
}

func mythical() int {
    var i int = 3
    defer fmt.Println("defer: ", i)

    i++
    return i
}
