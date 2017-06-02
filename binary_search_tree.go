package main

import "fmt"

type Node struct {
    right, left *Node
    value int
}

type BinarySearchTree struct {
    root *Node
}

func addNode(val int) *Node {
    return &Node{value: val}
}

func (tree *BinarySearchTree) Insert(val int) {
    if tree.root == nil {
        node := addNode(val)
        tree.root = node
    } else {
        tree.insert(val, tree.root)
    }
}

func (tree *BinarySearchTree) insert(val int, node *Node) *Node {
    if node == nil {
        node = addNode(val)
    } else if val <= node.value {
        node.left = tree.insert(val, node.left)
    } else if val > node.value {
        node.right = tree.insert(val, node.right)
    }
    return node
}

func preOrder(root *Node) {
    if root == nil {
        return
    }
    fmt.Println(root.value)
    preOrder(root.left)
    preOrder(root.right)
}

func main() {
    tree := new(BinarySearchTree)
    tree.Insert(10)
    tree.Insert(20)
    tree.Insert(30)
    tree.Insert(5)
    preOrder(tree.root)
}
