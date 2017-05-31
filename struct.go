package main

import "fmt"

type Node struct {
    val int
    next *Node
}

func main() {
    i := make(map[int]Node)
    j := make(map[int]Node)

    fmt.Println(i, j)
}
