package main

import (
	"fmt"
	"net/http"
)

func main() {
	response, err := http.Get("http://vipul.xyz/")
	fmt.Println(response, err)
}
