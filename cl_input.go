package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    var i int
    fmt.Print("Enter integer: ")
    fmt.Scan(&i)
    fmt.Println("You entered: ", i)

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, _ := reader.ReadString('\n')
    fmt.Print("You entered: ", text)
}
