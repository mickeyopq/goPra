package main

import (
	"fmt"
)

type person struct {
	name, age int
}

func (r *person) ff() int {
	// r.name += 100
	return r.name + r.age
}
func (r *person) ddd() int {
	// r.name += 100
	return r.name - r.age
}
func main() {
	fmt.Println("ffffff")
	fmt.Println(person{1234, 5555})
	fmt.Println(&person{9898, 4455})
	s := person{11, 22}
	fmt.Println(s.name)
	fmt.Println(s.ff())
	fmt.Println(s.name)
	fmt.Println(s.ddd())
}
