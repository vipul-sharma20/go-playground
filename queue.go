package main

import "fmt"

const size int = 10

type Queue struct {
    q [size]int
    front int
    rear int
}

func main() {
    queue := &Queue{front: -1, rear: -1}
    queue.push(12)
    fmt.Println("Pushed: 12")
    fmt.Println("Removed: ", queue.remove())
    queue.push(13)
    fmt.Println("Pushed: 13")
    queue.push(14)
    fmt.Println("Pushed: 14")
    fmt.Println("Removed: ", queue.remove())
    fmt.Println("Removed: ", queue.remove())
    fmt.Println("Removed: ", queue.remove())
}

func (queue *Queue) push(val int) {
    queue.rear++
    queue.q[queue.rear] = val
}

func (queue *Queue) remove() int {
    if queue.rear > queue.front {
        queue.front++
        item := queue.q[queue.front]
        return item
    } else {
        return -1
    }
}


