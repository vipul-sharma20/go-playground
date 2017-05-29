package main

import "fmt"

func wrapper() func() int {
	i := 0
	return func() int {
		i = i + 5
		return i
	}
}

func main() {
	f := wrapper()
	fmt.Println(f())
}
