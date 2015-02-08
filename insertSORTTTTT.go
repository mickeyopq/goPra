package main

import (
	"fmt"
)

func insertSort(x []int) []int {
	ll := len(x)
	k := 0
	for j := 0; j < ll-1; j++ {
		k = x[j+1]
		i := j
		for ; i >= 0 && x[i] > k; i-- {
			x[i+1] = x[i]
		}
		x[i+1] = k
	}
	return x
}
func main() {
	ss := []int{9, 42, 3, 4, 5, 1, 2}
	fmt.Println("=================")
	fmt.Println("before", ss)
	insertSort(ss)
	fmt.Println("after", ss)
}
