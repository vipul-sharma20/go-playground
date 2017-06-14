package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 1}
	b := [6]int{3, 4, 1, 2, 1, 3}

	var L [len(b)][len(a)]int

	for i := len(b) - 1; i >= 0; i-- {
		for j := len(a) - 1; j >= 0; j-- {
			if i == len(b)-1 || j == len(a)-1 {
				L[i][j] = 0
			} else if a[i] == b[j] {
				L[i][j] = L[i+1][j+1] + 1
			} else {
				L[i][j] = max(L[i+1][j], L[i][j+1])
			}
		}
	}
	fmt.Println(L)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
