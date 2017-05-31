package main

import "fmt"

const size int = 20

type Stack struct {
   s [size]int
   top int
}

func main() {
    stack := Stack{top: -1}
    stack.push(10)
    stack.push(20)
    fmt.Println(stack)
    fmt.Println(stack.pop())
    fmt.Println(stack.pop())
    fmt.Println(stack)
}

func (stack *Stack) push(val int) {
    stack.top++
    stack.s[stack.top] = val
}

func (stack *Stack) pop() int {
    item := stack.s[stack.top]
    stack.top--
    return item
}
