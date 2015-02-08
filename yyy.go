package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "fuck幹床前明天光"
	s2 := strings.Split(s1, "")
	fmt.Printf("s2,,,,%T\n", s2)

	var arr []string
	arr = append(arr, "一")
	arr = append(arr, "二")
	arr = append(arr, "三")
	arr = append(arr, "四")
	fmt.Println(arr[1])
	fmt.Println(arr)
	fmt.Printf("arr,,,,,%T\n", arr)
}
