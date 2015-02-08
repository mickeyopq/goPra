package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 34556
	c := float32(a)
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(c))
}
