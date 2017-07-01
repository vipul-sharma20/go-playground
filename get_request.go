package main

import (
    "io/ioutil"
    "net/http"
    "fmt"
)

type Test struct {
    name string
}

func catRequest() {
    request, _ := http.NewRequest("GET", "http://thecatapi.com/api/images/get", nil)

    query_string := request.URL.Query()
    query_string.Add("type", "jpg,png")
    query_string.Add("format", "html")
    query_string.Add("size", "med,full")
    request.URL.RawQuery = query_string.Encode()

    resp, _ := http.Get(request.URL.String())

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}

func main() {
    catRequest()
}
