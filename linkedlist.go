package main

import "fmt"

var ptr *Node

type Node struct {
    next *Node
    value int
}

type LinkedList struct {
    start *Node
}

func addNode(val int) *Node {
    return &Node{value: val}
}

func (ll *LinkedList) setNext(val int) {
    if ll.start == nil {
        node := addNode(val)
        ll.start = node
        ptr = node
    } else {
        fmt.Println(ll, ptr)
        ptr.next = addNode(val)
        ptr = ptr.next
    }
}

func main() {
    ll := new(LinkedList)
    ll.setNext(10)
    ll.setNext(20)
}
